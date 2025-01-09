package main

import (
	"circuit-breaker-services/src/handler"
	"circuit-breaker-services/src/service"
	"fmt"
	"net/http"
)

func main() {
	orderServiceURL := "http://order-service:8081/orders"
	svc := service.NewService(orderServiceURL)
	h := handler.NewHandler(svc)

	http.HandleFunc("/orders", h.HandleRequest)

	fmt.Println("Circuit Breaker Service running on port 8080")
	http.ListenAndServe(":8080", nil)
}
