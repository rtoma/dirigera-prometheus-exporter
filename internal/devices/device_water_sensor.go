package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_WATER_SENSOR_TYPE = "waterSensor"

type WaterSensorAttributes struct {
	BaseAttributes
	WaterLeakDetected bool `json:"waterLeakDetected"`
	BatteryPercentage int  `json:"batteryPercentage"`
}

func (a *WaterSensorAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(a.CustomName)

	labels := []string{a.CustomName}
	gaugeWaterSensorWaterLeakDetected.WithLabelValues(labels...).Set(boolToFloat64(a.WaterLeakDetected))
	gaugeWaterSensorBatteryPercentage.WithLabelValues(labels...).Set(float64(a.BatteryPercentage))
}

var (
	gaugeWaterSensorWaterLeakDetected = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_water_sensor_water_leak_detected",
			Help: "Water leak detected by IKEA BADRING Water Leakage Sensor (bool)",
		},
		[]string{"device_name"},
	)
	gaugeWaterSensorBatteryPercentage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_water_sensor_battery_percentage",
			Help: "Water sensor battery percentage",
		},
		[]string{"device_name"},
	)
)
