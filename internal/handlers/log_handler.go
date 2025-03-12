package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/models"
	"github.com/nikhildev/gofy/internal/services"
)

func LogRoutes(g *echo.Group) {
	g.GET("/:id", getLogHandler)
	g.POST("", saveLogHandler)
}

func getLogHandler(c echo.Context) error {
	id := c.Param("id")

	// Param validation should be done in the handler
	if id == "" {
		return c.String(http.StatusBadRequest, "ID is required")
	}
	res, _ := services.GetLog(id)
	return c.JSON(http.StatusOK, res)
}

func saveLogHandler(c echo.Context) error {
	// Here we declare a variable that will later be used to store the log message using the bing method
	var logMessage models.LogMessage

	message := c.Request().Body

	if message == nil {
		return c.String(http.StatusBadRequest, "Message is required")
	} else {
		// c.bind method is used to bind the request body to the logMessage variable based on the struct type of logMessage
		err := c.Bind(&logMessage)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "Invalid request")
		}

		err = services.CreateLog(logMessage.Message)
		if err != nil {
			log.Fatalln(err)
			return err
		}

		return nil
	}
}
