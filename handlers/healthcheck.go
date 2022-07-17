package handlers

import (
	"log"
	"net/http"

	"github.com/feynmaz/goecho/middlewares"
	"github.com/feynmaz/goecho/renderings"
	"github.com/labstack/echo/v4"
)

func HealthCheck(ctx echo.Context) error {
	if requestID := ctx.Get(middlewares.RequestIdContextKey); requestID != nil {
		log.Printf("RequestID: %s", requestID)
	}
	response := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return ctx.JSON(http.StatusOK, response)
}
