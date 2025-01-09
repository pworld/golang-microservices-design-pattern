package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	pb "api-gateway-service/generated"
)

func main() {
	fmt.Println("API Gateway started on port 8080...")

	// Debug: Print all environment variables
	fmt.Println("🔍 Loaded Environment Variables:")
	for _, e := range os.Environ() {
		fmt.Println(e)
	}

	// Check if running inside a Docker container
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		// Load .env file only in local environment
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading local.env file")
		}
	} else {
		fmt.Println("Running inside Docker, skipping .env file loading")
	}

	// Load JWT secret
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	} else {
		fmt.Println("JWT_SECRET successfully loaded")
	}

	// Initialize gRPC client for order service
	orderServiceAddr := "order-service:9090"
	orderClient := NewOrderServiceClient(orderServiceAddr)
	defer orderClient.Close()

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Gateway is running!")
	})

	// Protected routes (require authentication)
	protected := e.Group("")
	protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
	}))

	// 🛑 Allow login without authentication
	e.POST("/api/user/login", reverseProxy("user"))

	// User requests go via HTTP Reverse Proxy
	protected.Any("/api/user/*", reverseProxy("user"))

	// Order requests go via gRPC
	protected.GET("/api/order/:orderID", func(c echo.Context) error {
		orderID := c.Param("orderID")

		// Extract JWT token from Echo context
		userToken := c.Request().Header.Get("Authorization")
		if userToken == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing JWT Token"})
		}

		// Call the reusable gRPC function with JWT
		response, err := orderClient.CallOrderService(c.Request().Context(), "GetOrder", &pb.OrderRequest{OrderId: orderID}, userToken)
		if err != nil {
			log.Printf("Error calling Order Service: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve order"})
		}

		// Convert response to correct type
		orderResponse, ok := response.(*pb.OrderResponse)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid response type"})
		}

		return c.JSON(http.StatusOK, orderResponse)
	})

	// Start HTTP Server
	log.Fatal(e.Start(":8080"))
}
