package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_LIGHT_TYPE = "light"

type LightAttributes struct {
	BaseAttributes
	IsOn                bool `json:"isOn"`
	LightLevel          int  `json:"lightLevel"`
	ColorTemperature    int  `json:"colorTemperature"`
	ColorTemperatureMin int  `json:"colorTemperatureMin"`
	ColorTemperatureMax int  `json:"colorTemperatureMax"`
}

func (o *LightAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(o.CustomName)
	gaugeLightIsOn.WithLabelValues(o.CustomName).Set(boolToFloat64(o.IsOn))
	gaugeLightLevel.WithLabelValues(o.CustomName).Set(float64(o.LightLevel))
	gaugeLightColorTemperature.WithLabelValues(o.CustomName).Set(float64(o.ColorTemperature))
	gaugeLightColorTemperatureMin.WithLabelValues(o.CustomName).Set(float64(o.ColorTemperatureMin))
	gaugeLightColorTemperatureMax.WithLabelValues(o.CustomName).Set(float64(o.ColorTemperatureMax))
}

var (
	gaugeLightIsOn = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_is_on",
			Help: "Light is_on status",
		},
		[]string{"device_name"},
	)
	gaugeLightLevel = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_level",
			Help: "Light level",
		},
		[]string{"device_name"},
	)
	gaugeLightColorTemperature = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_color_temperature",
			Help: "Light color temperature",
		},
		[]string{"device_name"},
	)
	gaugeLightColorTemperatureMin = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_color_temperature_min",
			Help: "Light color temperature min value",
		},
		[]string{"device_name"},
	)
	gaugeLightColorTemperatureMax = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_color_temperature_max",
			Help: "Light color temperature max value",
		},
		[]string{"device_name"},
	)
)
