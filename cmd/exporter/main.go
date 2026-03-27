package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/collector"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/config"
	"github.com/rtoma/dirigera-prometheus-exporter/internal/events"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	metricsCollector := collector.NewMetricsCollector(cfg)
	if err := metricsCollector.CollectOnce(); err != nil {
		log.Fatal(err)
	}

	eventListener, err := events.NewWebsocketListener(cfg)
	if err != nil {
		log.Fatal(err)
	}

	go metricsCollector.StartIntervalCollector()
	go eventListener.Start()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	log.Printf("Starting HTTP server on %s", cfg.BindAddress)
	log.Fatal(http.ListenAndServe(cfg.BindAddress, nil))
}
