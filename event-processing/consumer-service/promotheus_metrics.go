package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics struct to store Prometheus metrics
type Metrics struct {
	EventsProduced prometheus.Counter
	EventsConsumed prometheus.Counter
}

// NewMetrics initializes Prometheus metrics
func NewMetrics(serviceName string) *Metrics {
	metrics := &Metrics{
		EventsProduced: prometheus.NewCounter(prometheus.CounterOpts{
			Name: serviceName + "_events_produced_total",
			Help: "Total number of events produced by the " + serviceName,
		}),
		EventsConsumed: prometheus.NewCounter(prometheus.CounterOpts{
			Name: serviceName + "_events_consumed_total",
			Help: "Total number of events consumed by the " + serviceName,
		}),
	}

	prometheus.MustRegister(metrics.EventsProduced)
	prometheus.MustRegister(metrics.EventsConsumed)

	return metrics
}

// StartMetricsServer starts an HTTP server to expose Prometheus metrics
func StartMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Printf("Prometheus metrics server running on port %s\n", port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatalf("Error starting Prometheus metrics server: %v", err)
		}
	}()
}
