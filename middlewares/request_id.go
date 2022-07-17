package middlewares

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

const (
	RequestIdContextKey = "request_id_context_key"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		requestID := uuid.NewV4()
		c.Logger().Infof("RequestID: %s", requestID)
		c.Set(RequestIdContextKey, requestID)
		return next(c)
	})
}
