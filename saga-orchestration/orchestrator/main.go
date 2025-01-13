package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "localhost:9092"
	topic       = "order_events"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		GroupID:  "saga_orchestrator",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		fmt.Printf("Received message: %s\n", string(msg.Value))
		handleSagaEvent(string(msg.Value))
	}
}

func handleSagaEvent(event string) {
	switch event {
	case "OrderCreated":
		fmt.Println("Processing order in Orchestrator...")
		publishEvent("ProcessOrder")
	case "PaymentProcessed":
		fmt.Println("Completing Order...")
		publishEvent("OrderCompleted")
	case "PaymentFailed":
		fmt.Println("Rolling back Order...")
		publishEvent("OrderCancelled")
	default:
		fmt.Println("Unknown event:", event)
	}
}

func publishEvent(event string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    "order_events",
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
