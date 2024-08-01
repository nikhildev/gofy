package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/db/mongodb"

	"github.com/nikhildev/gofy/internal/routes"
)

func main() {

	// This is where we call the mongodb.NewDbStore() function to create a new instance of the DbStore struct. This function returns a pointer to the DbStore struct and an error. We are ignoring the error for now, but in a real-world application, you should handle it properly.
	_, err := mongodb.NewDbStore(nil)
	if err != nil {
		panic(err)
	}

	//This creates a simple http echo server and starts it
	e := echo.New()

	// For simplicity, we are registering all the routes in the main function
	routes.RegisterRoutes(e)

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}
