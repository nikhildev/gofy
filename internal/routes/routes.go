package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nikhildev/gofy/internal/handlers"
)

// RegisterRoutes registers the routes for the application
func RegisterRoutes(e *echo.Echo) {
	// RequestID middleware generates a unique UUID for every request.
	// The UUID can be retrieved using echo.HeaderXRequestID header.
	// ex: requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	e.Use(middleware.RequestID())

	// Creating groups help us to create subroutes without repetitiveness.
	// Also, additional middleware can be applied to the group explicitly.
	healthGroup := e.Group("/health")
	handlers.HealthRoutes(healthGroup)

	userGroup := e.Group("/users")
	handlers.UserRoutes(userGroup)

}
