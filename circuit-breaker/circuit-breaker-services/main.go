package main

import (
	"circuit-breaker-services/src/handler"
	"circuit-breaker-services/src/service"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	orderServiceURL := "http://order-service:8081/orders"
	svc := service.NewService(orderServiceURL)
	h := handler.NewHandler(svc)

	http.HandleFunc("/orders", h.HandleRequest)
	http.Handle("/metrics", promhttp.Handler()) // Expose Prometheus metrics

	port := ":8080"
	fmt.Println("Circuit Breaker Service running on port", port)
	http.ListenAndServe(port, nil)
}
