package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
