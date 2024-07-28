package main

import (
	"github.com/labstack/echo/v4"

	"github.com/nikhildev/gofy/internal/routes"
)

func main() {

	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
