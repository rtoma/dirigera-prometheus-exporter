package devices

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const DEVICE_ELECTRICAL_SENSOR_TYPE = "electricalSensor"

type ElectricalAttributes struct {
	CurrentActivePower         *float64   `json:"currentActivePower"`
	CurrentAmps                *float64   `json:"currentAmps"`
	CurrentVoltage             *float64   `json:"currentVoltage"`
	CurrentAcFrequency         *float64   `json:"currentAcFrequency"`
	TotalEnergyConsumed        *float64   `json:"totalEnergyConsumed"`
	EnergyConsumedAtLastReset  *float64   `json:"energyConsumedAtLastReset"`
	TimeOfLastEnergyReset      *time.Time `json:"timeOfLastEnergyReset"`
	TotalEnergyConsumedUpdated *time.Time `json:"totalEnergyConsumedLastUpdated"`
}

type ElectricalSensorAttributes struct {
	BaseAttributes
	ElectricalAttributes
}

func (a *ElectricalSensorAttributes) CollectMetrics(device Device) {
	device.CollectMetrics(a.CustomName)
	a.collectElectricalMetrics(a.CustomName)
}

func (a *ElectricalAttributes) collectElectricalMetrics(deviceName string) {
	labels := []string{deviceName}
	if a.CurrentActivePower != nil {
		gaugeElectricalActivePowerWatts.WithLabelValues(labels...).Set(*a.CurrentActivePower)
	}
	if a.CurrentAmps != nil {
		gaugeElectricalCurrentAmperes.WithLabelValues(labels...).Set(*a.CurrentAmps)
	}
	if a.CurrentVoltage != nil {
		gaugeElectricalVoltageVolts.WithLabelValues(labels...).Set(*a.CurrentVoltage)
	}
	if a.CurrentAcFrequency != nil {
		gaugeElectricalFrequencyHertz.WithLabelValues(labels...).Set(*a.CurrentAcFrequency)
	}
	if a.TotalEnergyConsumed != nil {
		gaugeElectricalEnergyKilowattHours.WithLabelValues(labels...).Set(*a.TotalEnergyConsumed)
	}
	if a.EnergyConsumedAtLastReset != nil {
		gaugeElectricalEnergyAtResetKilowattHours.WithLabelValues(labels...).Set(*a.EnergyConsumedAtLastReset)
	}
	if a.TimeOfLastEnergyReset != nil {
		gaugeElectricalEnergyResetTimestamp.WithLabelValues(labels...).Set(float64(fixTimezone(*a.TimeOfLastEnergyReset).Unix()))
	}
	if a.TotalEnergyConsumedUpdated != nil {
		gaugeElectricalEnergyUpdatedTimestamp.WithLabelValues(labels...).Set(float64(fixTimezone(*a.TotalEnergyConsumedUpdated).Unix()))
	}
}

var (
	gaugeElectricalActivePowerWatts           = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_active_power_watts", Help: "Current active power reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalCurrentAmperes             = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_current_amperes", Help: "Current electrical current reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalVoltageVolts               = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_voltage_volts", Help: "Current electrical voltage reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalFrequencyHertz             = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_frequency_hertz", Help: "Current AC frequency reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalEnergyKilowattHours        = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_energy_kilowatt_hours", Help: "Cumulative total energy consumed reported by DIRIGERA (monotonically increasing since last reset; not a Prometheus counter)"}, []string{"device_name"})
	gaugeElectricalEnergyAtResetKilowattHours = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_energy_at_reset_kilowatt_hours", Help: "Cumulative energy consumed at the last reset reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalEnergyResetTimestamp       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_energy_reset_timestamp_seconds", Help: "Last energy reset timestamp reported by DIRIGERA"}, []string{"device_name"})
	gaugeElectricalEnergyUpdatedTimestamp     = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "ikea_dirigera_electrical_energy_updated_timestamp_seconds", Help: "Last total-energy update timestamp reported by DIRIGERA"}, []string{"device_name"})
)
