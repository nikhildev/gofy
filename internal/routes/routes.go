package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nikhildev/gofy/internal/handlers"
)

// RegisterRoutes registers the routes for the application
func RegisterRoutes(e *echo.Echo) {
	e.Use(middleware.RequestID())

	healthGroup := e.Group("/health")
	handlers.HealthRoutes(healthGroup)

	e.GET("/ping", handlers.PingHandler)
}
