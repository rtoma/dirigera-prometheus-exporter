package events

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/config"
)

type WebsocketListener struct {
	conn *websocket.Conn
	cfg  *config.DirigeraExporterConfig
}

type DirigeraEvent struct {
	Id          string          `json:"id"`
	Time        time.Time       `json:"time"`
	SpecVersion string          `json:"specversion"`
	Source      string          `json:"source"`
	Type        string          `json:"type"`
	Data        json.RawMessage `json:"data"`
}

func NewWebsocketListener(cfg *config.DirigeraExporterConfig) (*WebsocketListener, error) {

	conn, err := connectWebsocket(cfg)
	if err != nil {
		return nil, err
	}

	return &WebsocketListener{conn, cfg}, nil

}

func connectWebsocket(cfg *config.DirigeraExporterConfig) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, // TODO: do better
	}
	dialer.Proxy = http.ProxyFromEnvironment

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+cfg.AccessToken)

	url := strings.Replace(cfg.Url, "https://", "wss://", 1) // TODO: do better

	conn, _, err := dialer.Dial(url, headers)
	if err != nil {
		return nil, fmt.Errorf("error connecting to WebSocket: %v", err)
	}
	return conn, nil
}

func (d *WebsocketListener) Start() {
	log.Print("Starting websocket listener")

	for {
		_, message, err := d.conn.ReadMessage() // messageType always 1 = TextMessage
		if err != nil {
			counterDirigeraWebsocketReceiveFailures.WithLabelValues("read").Inc()
			log.Printf("Websocket ReadMessage error: %v", err)

			_ = d.conn.Close()

			// retry connection
			for {
				time.Sleep(time.Second) // to avoid crazy looping
				conn, err := connectWebsocket(d.cfg)
				if err == nil {
					d.conn = conn
					break
				}
				log.Printf("Error connecting to Websocket: %v", err)
			}
			continue
		}

		log.Printf("Received message: %s", message)

		var event DirigeraEvent
		if err := json.Unmarshal(message, &event); err != nil {
			counterDirigeraWebsocketReceiveFailures.WithLabelValues("decode").Inc()
			log.Printf("Error unmarshalling Websocket message %q: %v", message, err)
			continue
		}
		if err := processEvent(event); err != nil {
			log.Printf("Error processing event: %v", err)
		}
	}
}

func processEvent(event DirigeraEvent) error {
	// log.Printf("Received DirigeraEvent: %v", event)
	counterDirigeraWebsocketReceivedEvents.WithLabelValues(event.Source, event.Type).Inc()

	// switch event.Type {
	// case "":
	// }

	return nil
}

var (
	counterDirigeraWebsocketReceivedEvents = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ikea_dirigera_websocket_received_events",
			Help: "Received events from Websocket",
		},
		[]string{"source", "type"},
	)
	counterDirigeraWebsocketReceiveFailures = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ikea_dirigera_websocket_errors",
			Help: "Websocket errors",
		},
		[]string{"error"},
	)
)
