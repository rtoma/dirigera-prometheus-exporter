package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type ChildLockAttribute struct {
	ChildLock bool `json:"childLock"`
}

func (c *ChildLockAttribute) CollectMetrics(device Device, deviceName string) {
	gaugeDeviceChildLock.WithLabelValues(device.ID, device.DeviceType, deviceName).Set(boolToFloat64(c.ChildLock))
}

var (
	gaugeDeviceChildLock = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_device_child_lock_status",
			Help: "Device child lock status (bool)",
		},
		[]string{"device_id", "device_type", "device_name"},
	)
)
