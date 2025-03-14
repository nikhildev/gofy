package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthRoutes(g *echo.Group) {
	g.GET("", healthHandler)
	g.GET("/readiness", readinessHandler)
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func readinessHandler(c echo.Context) error {
	return c.String(http.StatusOK, "READY")
}
