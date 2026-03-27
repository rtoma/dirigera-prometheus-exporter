package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type StatusLightAttribute struct {
	StatusLight bool `json:"statusLight"`
}

func (a *StatusLightAttribute) CollectMetrics(device Device, deviceName string) {
	gaugeDeviceStatusLight.WithLabelValues(device.ID, device.DeviceType, deviceName).Set(boolToFloat64(a.StatusLight))
}

var (
	gaugeDeviceStatusLight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_device_status_light",
			Help: "Device status light status (bool)",
		},
		[]string{"device_id", "device_type", "device_name"},
	)
)
