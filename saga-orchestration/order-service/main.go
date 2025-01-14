package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "kafka:9092"
	topic       = "order_events"
)

func main() {
	fmt.Println("Order Service Started...")

	// Start HTTP Server
	http.HandleFunc("/create-order", createOrderHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// HTTP Handler to trigger order creation
func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received Order Creation Request")
	publishEvent("OrderCreated")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order Created Successfully!"))
}

// Publish Kafka Event
func publishEvent(event string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()
	err := w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(event)},
	)

	if err != nil {
		log.Fatalf("Failed to publish event: %v", err)
	}
	fmt.Printf("Published event: %s\n", event)
}
