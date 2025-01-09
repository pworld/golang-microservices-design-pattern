package routes

import (
	"user-service/handlers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes initializes user routes
func SetupRoutes(e *echo.Echo) {
	e.POST("/api/user/login", handlers.LoginUser) // Login route
	e.GET("/api/user/profile", handlers.GetProfile)
}
