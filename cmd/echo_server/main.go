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

	//This creates a simple http echo server and starts it
	server := echo.New()

	// go func() {
	log.Println("Registering routes")
	routes.RegisterRoutes(server)

	s := &http.Server{
		Addr:    ":3000",
		Handler: server,
	}
	// For simplicity, we are registering all the routes in the main function

	// Start the server
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("failed to start http server: %s", err)
	}

	return server

}
