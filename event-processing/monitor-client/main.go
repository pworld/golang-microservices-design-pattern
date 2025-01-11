package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "event-processing/proto" // Import gRPC protobuf

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEventServiceClient(conn)

	for {
		resp, err := client.GetProcessedMessages(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("Error fetching messages: %v", err)
		}

		fmt.Println("Processed Messages:")
		for _, msg := range resp.Messages {
			fmt.Printf("Message: %s at %s\n", msg.Content, msg.Timestamp)
		}
		time.Sleep(5 * time.Second)
	}
}
