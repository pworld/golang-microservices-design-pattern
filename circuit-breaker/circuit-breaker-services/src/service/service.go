package service

import (
	"circuit-breaker-services/src/circuitbreaker"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type Service interface {
	GetOrders() (string, error)
}

type serviceImpl struct {
	externalURL string
}

// Define Prometheus metrics
var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "circuit_breaker_requests_total",
			Help: "Total number of requests sent to order-service",
		},
		[]string{"status"},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(requestsTotal)
}

// NewService constructor
func NewService(url string) Service {
	return &serviceImpl{externalURL: url}
}

// GetOrders method with Prometheus monitoring
func (s *serviceImpl) GetOrders() (string, error) {
	response, err := circuitbreaker.CallOrderService(s.externalURL)
	if err != nil {
		fmt.Println("Circuit breaker triggered:", err)
		requestsTotal.WithLabelValues("failed").Inc()
		return "", err
	}

	requestsTotal.WithLabelValues("success").Inc()
	return response, nil
}
