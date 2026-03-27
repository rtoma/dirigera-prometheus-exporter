package devices

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_AIR_PURIFIER_TYPE = "airPurifier"

type AirPurifierAttributes struct {
	BaseAttributes
	ChildLockAttribute
	StatusLightAttribute
	FanMode           string  `json:"fanMode"`
	MotorState        int     `json:"motorState"`
	MotorRuntime      int     `json:"motorRuntime"`
	FilterElapsedTime int     `json:"filterElapsedTime"`
	FilterAlarmStatus bool    `json:"filterAlarmStatus"`
	FilterLifetime    int     `json:"filterLifetime"`
	CurrentPM25       float64 `json:"currentPM25"`
}

func (a *AirPurifierAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(a.CustomName)
	a.ChildLockAttribute.CollectMetrics(device, a.CustomName)
	a.StatusLightAttribute.CollectMetrics(device, a.CustomName)

	labels := []string{a.CustomName}
	gaugeAirPurifierPm25.WithLabelValues(labels...).Set(a.CurrentPM25)
	gaugeAirPurifierMotorState.WithLabelValues(labels...).Set(float64(a.MotorState))
	gaugeAirPurifierMotorRuntime.WithLabelValues(labels...).Set(float64(a.MotorRuntime))
	gaugeAirPurifierFilterElapsedTime.WithLabelValues(labels...).Set(float64(a.FilterElapsedTime))
	gaugeAirPurifierFilterLifetime.WithLabelValues(labels...).Set(float64(a.FilterLifetime))
	gaugeAirPurifierFilterAlarmStatus.WithLabelValues(labels...).Set(boolToFloat64(a.FilterAlarmStatus))
}

var (
	gaugeAirPurifierPm25 = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_pm25",
			Help: "AirPurifier PM2.5 level reported by IKEA Dirigera hub",
		},
		[]string{"device_name"},
	)
	gaugeAirPurifierMotorState = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_motor_state",
			Help: "AirPurifier motor state / speed",
		},
		[]string{"device_name"},
	)
	gaugeAirPurifierMotorRuntime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_motor_runtime",
			Help: "AirPurifier motor runtime",
		},
		[]string{"device_name"},
	)
	gaugeAirPurifierFilterElapsedTime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_filter_elapsed",
			Help: "AirPurifier filter elapsed time (minutes)",
		},
		[]string{"device_name"},
	)
	gaugeAirPurifierFilterLifetime = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_filter_lifetime",
			Help: "AirPurifier filter lifetime (minutes)",
		},
		[]string{"device_name"},
	)
	gaugeAirPurifierFilterAlarmStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ikea_dirigera_airpurifier_filter_alarm_status",
			Help: "AirPurifier filter alarm status (bool)",
		},
		[]string{"device_name"},
	)
)
