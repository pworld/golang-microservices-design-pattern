package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "order-service/generated"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// OrderServiceServer struct
type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

// GetOrder handles gRPC request and validates JWT
func (s *OrderServiceServer) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	// Extract JWT from metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	token := strings.TrimPrefix(authHeader[0], "Bearer ")
	if err := validateJWT(token); err != nil {
		return nil, fmt.Errorf("unauthorized: %v", err)
	}

	// Business logic
	return &pb.OrderResponse{
		OrderId: req.OrderId,
		Status:  "Processed",
	}, nil
}

// validateJWT decodes and verifies the JWT token
func validateJWT(tokenString string) error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// Parse and validate JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid or expired token")
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &OrderServiceServer{})

	log.Println("Order Service running on port 9090")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
