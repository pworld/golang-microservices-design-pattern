package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	metrics := NewMetrics("producer")
	StartMetricsServer("8080")

	topic := "events"

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Event #%d at %s", i, time.Now().Format(time.RFC3339))
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)

		if err == nil {
			metrics.EventsProduced.Inc()
			log.Printf("Produced message: %s", message)
		} else {
			log.Printf("Error producing message: %v", err)
		}

		time.Sleep(2 * time.Second)
	}
}
