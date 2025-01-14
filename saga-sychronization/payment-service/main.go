package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "kafka:9092"
	topic       = "order_events"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		GroupID:  "payment_service",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		fmt.Printf("Received message: %s\n", string(msg.Value))
		processPayment(string(msg.Value))
	}
}

func processPayment(event string) {
	if event == "ProcessOrder" {
		if rand.Intn(2) == 0 {
			fmt.Println("Payment Success")
			publishEvent("PaymentProcessed")
		} else {
			fmt.Println("Payment Failed")
			publishEvent("PaymentFailed")
		}
	}
}

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
