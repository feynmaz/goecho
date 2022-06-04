package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", index)
	e.Logger.Fatal(e.Start(":8080"))
}

func index(ctx echo.Context) error {
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Hello, World!"))
	return ctx.String(http.StatusOK, "Hello, World!")
}
