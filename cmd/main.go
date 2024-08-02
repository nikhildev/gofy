package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/db"
	"github.com/nikhildev/gofy/internal/routes"
)

func main() {

	// This is where we call the mongodb.NewStore() function to create a new instance of the Store struct. This function returns a pointer to the Store struct and an error. We are ignoring the error for now, but in a real-world application, you should handle it properly.
	_, err := db.NewStore(nil)
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
