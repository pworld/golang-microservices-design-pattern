package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"service-registry-service/src/discovery"
	"service-registry-service/src/routes"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("API Gateway started on port 8080...")

	// Check if running inside a Docker container
	if os.Getenv("RUNNING_IN_DOCKER") == "" {
		// Load .env file only in local environment
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatal("Error loading local.env file")
		}
	}

	// Load JWT secret from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20))) // Rate limiting

	// Public routes (no authentication)
	e.GET("/api/health", routes.HealthCheck)

	// Protected routes (require authentication)
	protected := e.Group("")
	protected.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
	}))

	protected.Any("/api/user/*", reverseProxy("user-service"))
	protected.Any("/api/order/*", reverseProxy("order-service"))

	// Start HTTP Server
	log.Fatal(e.Start(":8080"))
}

func reverseProxy(serviceName string) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceAddr, err := discovery.DiscoverService(serviceName)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Service not found"})
		}

		targetURL := fmt.Sprintf("http://%s%s", serviceAddr, c.Request().URL.Path)
		resp, err := http.Get(targetURL)
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Service unavailable"})
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		return c.JSONBlob(resp.StatusCode, body)
	}
}
