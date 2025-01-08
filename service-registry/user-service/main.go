package main

import (
	"fmt"
	"log"

	"user-service/consul"
	// "user-service/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("User Service started on port 8081...")

	// Register service with Consul
	consul.RegisterService("user-service", 8081)

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// routes.SetupRoutes(e)

	// Start HTTP Server
	log.Fatal(e.Start(":8081"))
}
