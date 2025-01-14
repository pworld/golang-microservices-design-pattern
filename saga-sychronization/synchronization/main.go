package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaBroker = os.Getenv("KAFKA_BROKER")

const topic = "saga_synchronization"

type SyncTransactionRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

func main() {
	if kafkaBroker == "" {
		log.Fatal("❌ KAFKA_BROKER environment variable is not set")
	}

	fmt.Println("✅ Saga Synchronization Service Started on port 8083...")
	fmt.Printf("✅ Connecting to Kafka at: %s\n", kafkaBroker)

	// Retry Kafka connection until the topic is available
	for i := 0; i < 5; i++ {
		if topicExists(topic) {
			break
		}
		fmt.Println("🔄 Waiting for Kafka topic to be available...")
		time.Sleep(5 * time.Second)
	}

	// Start HTTP Server
	go func() {
		http.HandleFunc("/sync-transaction", syncTransactionHandler)
		log.Fatal(http.ListenAndServe(":8083", nil))
	}()

	select {} // Keep service running
}

// Handle Synchronous Transaction Request
func syncTransactionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request SyncTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("🔄 Processing Synchronous Transaction for Order ID: %s\n", request.OrderID)

	// Simulating synchronous process with a delay
	time.Sleep(2 * time.Second)

	// Publish Event to Kafka
	publishEvent("SyncTransactionCompleted", request.OrderID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Synchronous transaction completed successfully"))
}

// Check if Kafka topic exists
func topicExists(topic string) bool {
	conn, err := kafka.Dial("tcp", kafkaBroker)
	if err != nil {
		fmt.Println("❌ Error connecting to Kafka:", err)
		return false
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions(topic)
	if err != nil {
		return false
	}
	return len(partitions) > 0
}

// Publish Kafka Event
func publishEvent(eventType, orderID string) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()
	eventData := fmt.Sprintf("%s for Order ID: %s", eventType, orderID)
	err := w.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(eventData)},
	)

	if err != nil {
		log.Fatalf("❌ Failed to publish event: %v", err)
	}
	fmt.Printf("✅ Published event: %s\n", eventData)
}
