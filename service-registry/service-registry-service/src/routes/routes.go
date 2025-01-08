package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthCheck endpoint
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "API Gateway is running!")
}
