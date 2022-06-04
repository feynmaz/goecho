package handlers

import (
	"net/http"

	"github.com/feynmaz/goecho/renderings"
	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {
	response := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return ctx.JSON(http.StatusOK, response)
}
