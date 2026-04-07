package collector

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/config"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/devices"
)

type MetricsCollector struct {
	cfg *config.DirigeraExporterConfig
}

func NewMetricsCollector(cfg *config.DirigeraExporterConfig) *MetricsCollector {
	log.Printf("Scraping %s, every %s", cfg.Url, cfg.Interval)
	return &MetricsCollector{cfg}
}

func (c *MetricsCollector) StartIntervalCollector() {

	ticker := time.NewTicker(c.cfg.Interval)
	defer ticker.Stop()

	for range ticker.C {
		if err := c.CollectOnce(); err != nil {
			log.Printf("CollectOnce error: %v", err)
		}
	}
}

// CollectOnce is a helper method to collect metrics once
func (c *MetricsCollector) CollectOnce() error {

	start := time.Now()

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // TODO: do better
			Proxy:           http.ProxyFromEnvironment,
		},
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", c.cfg.Url+"/v1/devices", nil)
	if err != nil {
		counterDirigeraScrapeFailures.WithLabelValues("prepare-request").Inc()
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.cfg.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		counterDirigeraScrapeFailures.WithLabelValues("request-failed").Inc()
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var deviceList []devices.Device
	var buf bytes.Buffer
	if err := json.NewDecoder(io.TeeReader(resp.Body, &buf)).Decode(&deviceList); err != nil {
		log.Printf("Received unexpected response from /v1/devices request: data=%s", buf.String())
		counterDirigeraScrapeFailures.WithLabelValues("decode-response").Inc()
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, device := range deviceList {
		switch device.DeviceType {
		case devices.DEVICE_AIR_PURIFIER_TYPE:
			var airPurifier devices.AirPurifierAttributes
			if err := json.Unmarshal(device.Attributes, &airPurifier); err != nil {
				return fmt.Errorf("error unmarshalling AirPurifierAttributes: %v", err)
			}
			airPurifier.CollectMetrics(device)

		case devices.DEVICE_ENVIRONMENT_SENSOR_TYPE:
			var envSensor devices.EnvironmentSensorAttributes
			if err := json.Unmarshal(device.Attributes, &envSensor); err != nil {
				return fmt.Errorf("error unmarshalling EnvironmentSensorAttributes: %v", err)
			}
			envSensor.CollectMetrics(device)

		case devices.DEVICE_GATEWAY_TYPE:
			var gw devices.GatewayAttributes
			if err := json.Unmarshal(device.Attributes, &gw); err != nil {
				return fmt.Errorf("error unmarshalling GatewayAttributes: %v", err)
			}
			gw.CollectMetrics(device)

		case devices.DEVICE_OUTLET_TYPE:
			var outlet devices.OutletAttributes
			if err := json.Unmarshal(device.Attributes, &outlet); err != nil {
				return fmt.Errorf("error unmarshalling OutletAttributes: %v", err)
			}
			outlet.CollectMetrics(device)

		case devices.DEVICE_SPEAKER_TYPE:
			var speaker devices.SpeakerAttributes
			if err := json.Unmarshal(device.Attributes, &speaker); err != nil {
				return fmt.Errorf("error unmarshalling SpeakerAttributes: %v", err)
			}
			speaker.CollectMetrics(device)

		case devices.DEVICE_LIGHT_TYPE:
			var light devices.LightAttributes
			if err := json.Unmarshal(device.Attributes, &light); err != nil {
				return fmt.Errorf("error unmarshalling LightAttributes: %v", err)
			}
			light.CollectMetrics(device)

		case devices.DEVICE_LIGHT_CONTROLLER_TYPE:
			var controller devices.LightControllerAttributes
			if err := json.Unmarshal(device.Attributes, &controller); err != nil {
				return fmt.Errorf("error unmarshalling LightControllerAttributes: %v", err)
			}
			controller.CollectMetrics(device)

		case devices.DEVICE_WATER_SENSOR_TYPE:
			var sensor devices.WaterSensorAttributes
			if err := json.Unmarshal(device.Attributes, &sensor); err != nil {
				return fmt.Errorf("error unmarshalling WaterSensorAttributes: %v", err)
			}
			sensor.CollectMetrics(device)

		default:
			device.CollectMetrics("")
			// log.Printf("Unknown device, type=%s, deviceType=%s, id=%s, lastSeen=%s, reachable=%v, attributes=%s",
			// 	device.Type, device.DeviceType, device.ID, device.LastSeen, device.IsReachable, device.Attributes)
			log.Printf("Unknown device: %+v", device)
		}
	}

	histogramDirigeraScrapeDuration.Observe(time.Since(start).Seconds())

	return nil
}

var (
	histogramDirigeraScrapeDuration = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name: "ikea_dirigera_scrape_duration",
			Help: "Time it takes to succesfully scrape the Dirigera (seconds)",
		},
	)
	counterDirigeraScrapeFailures = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ikea_dirigera_scrape_errors",
			Help: "Scrape errors",
		},
		[]string{"error"},
	)
)
