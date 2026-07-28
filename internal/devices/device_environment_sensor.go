package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_ENVIRONMENT_SENSOR_TYPE = "environmentSensor"

type EnvironmentSensorAttributes struct {
	BaseAttributes
	Temperature     float64  `json:"currentTemperature"`
	Humidity        int      `json:"currentRH"`
	PM25            int      `json:"currentPM25"`
	MaxMeasuredPM25 int      `json:"maxMeasuredPM25"`
	MinMeasuredPM25 int      `json:"minMeasuredPM25"`
	VocIndex        int      `json:"vocIndex"`
	CO2             *float64 `json:"currentCO2"`
	MaxMeasuredCO2  *float64 `json:"maxMeasuredCO2"`
	MinMeasuredCO2  *float64 `json:"minMeasuredCO2"`
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
	if a.CO2 != nil {
		gaugeEnvironmentSensorCO2PPM.WithLabelValues(labels...).Set(*a.CO2)
	}
	if a.MaxMeasuredCO2 != nil {
		gaugeEnvironmentSensorMaxMeasuredCO2PPM.WithLabelValues(labels...).Set(*a.MaxMeasuredCO2)
	}
	if a.MinMeasuredCO2 != nil {
		gaugeEnvironmentSensorMinMeasuredCO2PPM.WithLabelValues(labels...).Set(*a.MinMeasuredCO2)
	}
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
	gaugeEnvironmentSensorCO2PPM = promauto.NewGaugeVec(
		prometheus.GaugeOpts{Name: "ikea_dirigera_environment_sensor_co2_ppm", Help: "Current CO2 concentration reported by IKEA Dirigera hub"},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorMaxMeasuredCO2PPM = promauto.NewGaugeVec(
		prometheus.GaugeOpts{Name: "ikea_dirigera_environment_sensor_max_co2_ppm", Help: "Maximum measured CO2 concentration reported by IKEA Dirigera hub"},
		[]string{"device_name"},
	)
	gaugeEnvironmentSensorMinMeasuredCO2PPM = promauto.NewGaugeVec(
		prometheus.GaugeOpts{Name: "ikea_dirigera_environment_sensor_min_co2_ppm", Help: "Minimum measured CO2 concentration reported by IKEA Dirigera hub"},
		[]string{"device_name"},
	)
)
