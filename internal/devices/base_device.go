package devices

import (
	"encoding/json"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Device struct {
	ID          string          `json:"id"`
	Type        string          `json:"type"`
	DeviceType  string          `json:"deviceType"`
	IsReachable bool            `json:"isReachable"`
	CreatedAt   time.Time       `json:"createdAt"`
	LastSeen    time.Time       `json:"lastSeen"`
	Attributes  json.RawMessage `json:"attributes"`
}

func (d *Device) CollectMetrics(deviceName string) {
	labels := []string{d.ID, d.DeviceType, deviceName}

	gaugeDeviceLastSeen.WithLabelValues(labels...).Set(float64(fixTimezone(d.LastSeen).Unix()))
	gaugeDeviceIsReachable.WithLabelValues(labels...).Set(boolToFloat64(d.IsReachable))
}

var (
	gaugeDeviceIsReachable = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_device_reachable",
			Help: "Device is reachable",
		},
		[]string{"device_id", "device_type", "device_name"},
	)

	gaugeDeviceLastSeen = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_device_last_seen",
			Help: "AirPurifier last seen (unix timestamp in seconds)",
		},
		[]string{"device_id", "device_type", "device_name"},
	)
)
