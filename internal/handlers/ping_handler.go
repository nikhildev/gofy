package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PingHandler(c echo.Context) error {
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	fmt.Println("RequestID: ", requestId)
	return c.String(http.StatusOK, fmt.Sprintf("pong: {RequestID: %s}", requestId))
}
