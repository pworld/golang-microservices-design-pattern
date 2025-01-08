package main

import (
	"fmt"
	"log"

	"order-service/consul"
	// "order-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Order Service started on port 8082...")

	// Register service with Consul
	consul.RegisterService("order-service", 8082)

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// routes.SetupRoutes(e)

	// Start HTTP Server
	log.Fatal(e.Start(":8082"))
}
