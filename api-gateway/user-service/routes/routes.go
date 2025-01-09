package routes

import (
	"user-service/handlers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes initializes user routes
func SetupRoutes(e *echo.Echo) {
	e.POST("/login", handlers.LoginUser)
	e.GET("/profile", handlers.GetProfile)
}
