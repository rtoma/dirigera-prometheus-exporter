package devices

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestElectricalSensorAttributes_Unmarshal(t *testing.T) {
	tests := []struct {
		name            string
		payload         string
		wantActivePower *float64
		wantTotalEnergy *float64
		wantResetTime   bool
		wantUpdatedTime bool
	}{
		{
			name: "all fields present",
			payload: `{
				"customName": "Full Sensor",
				"currentActivePower": 123.4,
				"currentAmps": 0.56,
				"currentVoltage": 230.1,
				"currentAcFrequency": 50.0,
				"totalEnergyConsumed": 999.9,
				"energyConsumedAtLastReset": 100.0,
				"timeOfLastEnergyReset": "2026-01-01T00:00:00Z",
				"totalEnergyConsumedLastUpdated": "2026-01-02T00:00:00Z"
			}`,
			wantActivePower: floatPtr(123.4),
			wantTotalEnergy: floatPtr(999.9),
			wantResetTime:   true,
			wantUpdatedTime: true,
		},
		{
			name: "optional fields absent",
			payload: `{
				"customName": "Partial Sensor"
			}`,
			wantActivePower: nil,
			wantTotalEnergy: nil,
			wantResetTime:   false,
			wantUpdatedTime: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var attrs ElectricalSensorAttributes
			if err := json.Unmarshal([]byte(tt.payload), &attrs); err != nil {
				t.Fatalf("unexpected unmarshal error: %v", err)
			}

			assertFloatPtrEqual(t, "CurrentActivePower", attrs.CurrentActivePower, tt.wantActivePower)
			assertFloatPtrEqual(t, "TotalEnergyConsumed", attrs.TotalEnergyConsumed, tt.wantTotalEnergy)

			if (attrs.TimeOfLastEnergyReset != nil) != tt.wantResetTime {
				t.Errorf("TimeOfLastEnergyReset presence = %v, want %v", attrs.TimeOfLastEnergyReset != nil, tt.wantResetTime)
			}
			if (attrs.TotalEnergyConsumedUpdated != nil) != tt.wantUpdatedTime {
				t.Errorf("TotalEnergyConsumedUpdated presence = %v, want %v", attrs.TotalEnergyConsumedUpdated != nil, tt.wantUpdatedTime)
			}
		})
	}
}

func TestElectricalAttributes_collectElectricalMetrics(t *testing.T) {
	t.Run("present fields are exported with correct values", func(t *testing.T) {
		attrs := ElectricalAttributes{
			CurrentActivePower:  floatPtr(42.5),
			CurrentAmps:         floatPtr(1.5),
			TotalEnergyConsumed: floatPtr(777.7),
		}

		resetTime := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
		attrs.TimeOfLastEnergyReset = &resetTime

		deviceName := "electrical-collect-present"
		attrs.collectElectricalMetrics(deviceName)

		if got := testutil.ToFloat64(gaugeElectricalActivePowerWatts.WithLabelValues(deviceName)); got != 42.5 {
			t.Errorf("gaugeElectricalActivePowerWatts = %v, want 42.5", got)
		}
		if got := testutil.ToFloat64(gaugeElectricalCurrentAmperes.WithLabelValues(deviceName)); got != 1.5 {
			t.Errorf("gaugeElectricalCurrentAmperes = %v, want 1.5", got)
		}
		if got := testutil.ToFloat64(gaugeElectricalEnergyKilowattHours.WithLabelValues(deviceName)); got != 777.7 {
			t.Errorf("gaugeElectricalEnergyKilowattHours = %v, want 777.7", got)
		}

		wantResetTimestamp := float64(fixTimezone(resetTime).Unix())
		if got := testutil.ToFloat64(gaugeElectricalEnergyResetTimestamp.WithLabelValues(deviceName)); got != wantResetTimestamp {
			t.Errorf("gaugeElectricalEnergyResetTimestamp = %v, want %v", got, wantResetTimestamp)
		}
	})

	t.Run("nil optional fields do not panic", func(t *testing.T) {
		attrs := ElectricalAttributes{}

		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("collectElectricalMetrics panicked on nil fields: %v", r)
			}
		}()

		attrs.collectElectricalMetrics("electrical-collect-nil")
	})
}

func floatPtr(v float64) *float64 {
	return &v
}

func assertFloatPtrEqual(t *testing.T, field string, got, want *float64) {
	t.Helper()
	if (got == nil) != (want == nil) {
		t.Errorf("%s presence = %v, want %v", field, got != nil, want != nil)
		return
	}
	if got != nil && *got != *want {
		t.Errorf("%s = %v, want %v", field, *got, *want)
	}
}
