package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_LIGHT_CONTROLLER_TYPE = "lightController"

type LightControllerAttributes struct {
	BaseAttributes
	IsOn              bool `json:"isOn"`
	LightLevel        int  `json:"lightLevel"`
	BatteryPercentage int  `json:"batteryPercentage"`
}

func (o *LightControllerAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(o.CustomName)
	gaugeLightControllerIsOn.WithLabelValues(o.CustomName).Set(boolToFloat64(o.IsOn))
	gaugeLightControllerLightLevel.WithLabelValues(o.CustomName).Set(float64(o.LightLevel))
	gaugeLightControllerBatteryPercentage.WithLabelValues(o.CustomName).Set(float64(o.BatteryPercentage))
}

var (
	gaugeLightControllerIsOn = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_controller_is_on",
			Help: "Light controller is_on status",
		},
		[]string{"device_name"},
	)
	gaugeLightControllerLightLevel = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_controller_light_level",
			Help: "Light controller light level",
		},
		[]string{"device_name"},
	)
	gaugeLightControllerBatteryPercentage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_light_controller_battery_percentage",
			Help: "Light controller battery percentage",
		},
		[]string{"device_name"},
	)
)
