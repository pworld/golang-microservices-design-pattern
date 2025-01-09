package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "api-gateway-service/generated"
)

// OrderServiceClient wraps a gRPC client for OrderService.
type OrderServiceClient struct {
	client pb.OrderServiceClient
	conn   *grpc.ClientConn
}

var (
	instance *OrderServiceClient
	once     sync.Once
)

// NewOrderServiceClient initializes a gRPC client connection only once (Singleton).
func NewOrderServiceClient(orderServiceAddr string) *OrderServiceClient {
	once.Do(func() {
		// Establish a new gRPC client connection
		conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Printf("Failed to create gRPC client: %v\n", err)
			return
		}

		// Assign instance once
		instance = &OrderServiceClient{
			client: pb.NewOrderServiceClient(conn),
			conn:   conn,
		}
	})

	return instance
}

// CallOrderService is a generic function to call different gRPC methods.
func (o *OrderServiceClient) CallOrderService(ctx context.Context, method string, request interface{}, token string) (interface{}, error) {
	// Add JWT token to gRPC metadata
	md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", token))
	ctx = metadata.NewOutgoingContext(ctx, md)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch method {
	case "GetOrder":
		req, ok := request.(*pb.OrderRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type for GetOrder")
		}
		return o.client.GetOrder(ctx, req)

	default:
		return nil, fmt.Errorf("unsupported gRPC method: %s", method)
	}
}

// Close gracefully closes the gRPC connection.
func (o *OrderServiceClient) Close() {
	if o.conn != nil {
		o.conn.Close()
	}
}
