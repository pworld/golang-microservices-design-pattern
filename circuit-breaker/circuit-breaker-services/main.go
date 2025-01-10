package main

import (
	"circuit-breaker-services/src/handler"
	"circuit-breaker-services/src/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting Circuit Breaker Service")

	orderServiceURL := "http://order-service:8081/orders"
	svc := service.NewService(orderServiceURL)
	h := handler.NewHandler(svc)

	http.HandleFunc("/orders", h.HandleRequest)

	port := ":8080"
	logrus.Infof("Circuit Breaker Service running on port %s", port)
	http.ListenAndServe(port, nil)
}
