package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	userRepo.Add(repositories.User{
		ID:   1,
		Name: "John Doe",
	})
	return c.String(http.StatusOK, "ADDED")
}
