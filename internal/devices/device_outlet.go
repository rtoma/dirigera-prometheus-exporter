package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_OUTLET_TYPE = "outlet"

type OutletAttributes struct {
	BaseAttributes
	ChildLockAttribute
	StatusLightAttribute
	IsOn       bool `json:"isOn"`
	LightLevel int  `json:"lightLevel"`
}

func (o *OutletAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(o.CustomName)
	o.ChildLockAttribute.CollectMetrics(device, o.CustomName)
	o.StatusLightAttribute.CollectMetrics(device, o.CustomName)
	gaugeOutletIsOn.WithLabelValues(o.CustomName).Set(boolToFloat64(o.IsOn))
	gaugeOutletLightLevel.WithLabelValues(o.CustomName).Set(float64(o.LightLevel))
}

var (
	gaugeOutletIsOn = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_outlet_is_on",
			Help: "Outlet is_on status",
		},
		[]string{"device_name"},
	)
	gaugeOutletLightLevel = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_outlet_light_level",
			Help: "Outlet light level",
		},
		[]string{"device_name"},
	)
)
