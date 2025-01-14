package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/segmentio/kafka-go"
)

// Read Kafka broker from environment variable
var kafkaBroker = os.Getenv("KAFKA_BROKER")

const topic = "order_events"

type OrderRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

func main() {
	if kafkaBroker == "" {
		log.Fatal("KAFKA_BROKER environment variable is not set")
	}

	fmt.Println("Saga Orchestrator Started on port 8080...")
	fmt.Printf("Connecting to Kafka at: %s\n", kafkaBroker)

	// Start HTTP Server
	go func() {
		http.HandleFunc("/create-order", createOrderHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	select {} // Keep service running
}

// Handle Order Creation Request
func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var order OrderRequest
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received order creation request: %+v\n", order)
	publishEvent("OrderCreated")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Order Created and Saga started!"))
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
