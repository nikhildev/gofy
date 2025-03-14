package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nikhildev/gofy/internal/models"
	"github.com/nikhildev/gofy/internal/repositories"
)

func UserRoutes(g *echo.Group) {
	g.GET("", getAllHandler)
	g.GET("/{id}", getByIdHandler)
	g.POST("", addHandler)
}

var userRepo = repositories.NewUserRepository()

func getAllHandler(c echo.Context) error {
	users := userRepo.GetAll()
	return c.JSON(http.StatusOK, users)
}

func getByIdHandler(c echo.Context) error {
	return c.String(http.StatusOK, "READY")
}

func addHandler(c echo.Context) error {
	payload := c.Request().Body

	if payload == nil {
		return c.String(http.StatusBadRequest, "Payload is required")
	}

	var user models.User
	var err error
	err = c.Bind(&user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	_, err = userRepo.Add(user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
