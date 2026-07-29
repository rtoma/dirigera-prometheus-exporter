package devices

import (
	"encoding/json"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	dto "github.com/prometheus/client_model/go"
)

func TestEnvironmentSensorAttributes_CO2Unmarshal(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		wantCO2    *float64
		wantMaxCO2 *float64
		wantMinCO2 *float64
	}{
		{
			name: "co2 fields present",
			payload: `{
				"customName": "CO2 Sensor",
				"currentCO2": 450.0,
				"maxMeasuredCO2": 900.0,
				"minMeasuredCO2": 400.0
			}`,
			wantCO2:    floatPtr(450.0),
			wantMaxCO2: floatPtr(900.0),
			wantMinCO2: floatPtr(400.0),
		},
		{
			name: "co2 fields absent stay nil",
			payload: `{
				"customName": "No CO2 Sensor"
			}`,
			wantCO2:    nil,
			wantMaxCO2: nil,
			wantMinCO2: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attrs EnvironmentSensorAttributes
			if err := json.Unmarshal([]byte(tt.payload), &attrs); err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}

			assertFloatPtrEqual(t, "CO2", attrs.CO2, tt.wantCO2)
			assertFloatPtrEqual(t, "MaxMeasuredCO2", attrs.MaxMeasuredCO2, tt.wantMaxCO2)
			assertFloatPtrEqual(t, "MinMeasuredCO2", attrs.MinMeasuredCO2, tt.wantMinCO2)
		})
	}
}

func TestEnvironmentSensorAttributes_CollectMetrics(t *testing.T) {
	device := Device{ID: "env-sensor-1", DeviceType: DEVICE_ENVIRONMENT_SENSOR_TYPE}

	t.Run("co2 present is exported with correct value", func(t *testing.T) {
		attrs := EnvironmentSensorAttributes{
			BaseAttributes: BaseAttributes{CustomName: "env-collect-present"},
			CO2:            floatPtr(600.0),
		}

		attrs.CollectMetrics(device)

		if got := testutil.ToFloat64(gaugeEnvironmentSensorCO2PPM.WithLabelValues("env-collect-present")); got != 600.0 {
			t.Errorf("gaugeEnvironmentSensorCO2PPM = %v, want 600.0", got)
		}
	})

	t.Run("nil co2 does not panic and is not exported", func(t *testing.T) {
		attrs := EnvironmentSensorAttributes{
			BaseAttributes: BaseAttributes{CustomName: "env-collect-nil"},
		}

		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("CollectMetrics panicked on nil CO2 field: %v", r)
			}
		}()

		attrs.CollectMetrics(device)

		if hasSeriesWithLabelValue(t, gaugeEnvironmentSensorCO2PPM, "device_name", "env-collect-nil") {
			t.Errorf("expected no CO2 series to be exported for a device with nil CO2")
		}
	})
}

func hasSeriesWithLabelValue(t *testing.T, c prometheus.Collector, labelName, labelValue string) bool {
	t.Helper()

	ch := make(chan prometheus.Metric)
	go func() {
		c.Collect(ch)
		close(ch)
	}()

	for m := range ch {
		var pb dto.Metric
		if err := m.Write(&pb); err != nil {
			t.Fatalf("failed to write metric: %v", err)
		}
		for _, l := range pb.GetLabel() {
			if l.GetName() == labelName && l.GetValue() == labelValue {
				return true
			}
		}
	}
	return false
}
