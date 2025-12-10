package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/routes"
)

func main() {
	startApiServer()
}

func startApiServer() *echo.Echo {
	log.Println("Starting API server")

	e := echo.New()

	log.Println("Registering routes")
	routes.RegisterRoutes(e)

	s := newHTTPServer(e)

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("failed to start http server: %s", err)
	}

	return e
}

// newHTTPServer constructs a net/http Server that uses the provided Echo instance as the handler.
func newHTTPServer(e *echo.Echo) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: e,
	}
}
