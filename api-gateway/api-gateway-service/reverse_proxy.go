package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

// Define service addresses
var serviceMap = map[string]string{
	"user":  "http://user-service:8081",
	"order": "http://order-service:9090",
}

// Reverse proxy function
func reverseProxy(service string) echo.HandlerFunc {
	return func(c echo.Context) error {
		serviceAddr, exists := serviceMap[service]
		if !exists {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Service not found"})
		}

		target, err := url.Parse(serviceAddr)
		if err != nil {
			log.Printf("Failed to parse service URL: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid service address"})
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(target)

		// Ensure request path is correctly rewritten
		c.Request().URL.Path = strings.TrimPrefix(c.Request().URL.Path, "/api/user")

		// Forward request
		proxy.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
