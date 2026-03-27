package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_ENVIRONMENT_SENSOR_TYPE = "environmentSensor"

type EnvironmentSensorAttributes struct {
	BaseAttributes
	Temperature     int `json:"currentTemperature"`
	Humidity        int `json:"currentRH"`
	PM25            int `json:"currentPM25"`
	MaxMeasuredPM25 int `json:"maxMeasuredPM25"`
	MinMeasuredPM25 int `json:"minMeasuredPM25"`
	VocIndex        int `json:"vocIndex"`
}

func (a *EnvironmentSensorAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(a.CustomName)

	labels := []string{a.CustomName}
	gaugeEnvironmentSensorTemperature.WithLabelValues(labels...).Set(float64(a.Temperature))
	gaugeEnvironmentSensorHumidity.WithLabelValues(labels...).Set(float64(a.Humidity))
	gaugeEnvironmentSensorPM25.WithLabelValues(labels...).Set(float64(a.PM25))
	gaugeEnvironmentSensorMaxMeasuredPM25.WithLabelValues(labels...).Set(float64(a.MaxMeasuredPM25))
	gaugeEnvironmentSensorMinMeasuredPM25.WithLabelValues(labels...).Set(float64(a.MinMeasuredPM25))
	gaugeEnvironmentSensorVocIndex.WithLabelValues(labels...).Set(float64(a.VocIndex))
}

var (
	gaugeEnvironmentSensorTemperature = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_temperature",
			Help: "Environment Sensor Temperature reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorHumidity = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_humidity",
			Help: "Environment Sensor Humidity reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorPM25 = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_pm25",
			Help: "Environment Sensor PM25 reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorMaxMeasuredPM25 = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_max_pm25",
			Help: "Environment Sensor Max Measured PM25 reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorMinMeasuredPM25 = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_min_pm25",
			Help: "Environment Sensor Min Measured PM25 reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorVocIndex = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_environment_sensor_voc_index",
			Help: "Environment Sensor VOC Index reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
)
