package main

import (
	"github.com/labstack/echo/v4"

	"github.com/nikhildev/gofy/internal/routes"
)

func main() {

	//This creates a simple http echo server and starts it
	e := echo.New()

	// For simplicity, we are registering all the routes in the main function
	routes.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}
