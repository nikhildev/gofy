package handlers

import (
	"github.com/labstack/echo/v4"
)

func CreateIssueHandler(e echo.Context) error {
	return e.String(200, "Create Issue Handler")
}
