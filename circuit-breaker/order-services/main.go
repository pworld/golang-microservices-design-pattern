package main

import (
	"fmt"
	"net/http"
	"order-services/src/service"
)

func main() {
	http.HandleFunc("/orders", service.GetOrderHandler)

	fmt.Println("Order Service running on port 8081")
	http.ListenAndServe(":8081", nil)
}
