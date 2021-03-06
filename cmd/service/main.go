package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/feynmaz/goecho/handlers"
	"github.com/feynmaz/goecho/middlewares"
	"github.com/feynmaz/goecho/models"
)

func main() {
	e := echo.New()

	// Signing Key for the auth middleware
	var signingKey = []byte("secret")
	e.Pre(middlewares.RequestIDMiddleware)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	// Add DB to context
	db, err := sql.Open("sqlite3", "./service.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.DBContextKey, db)
			return next(c)
		}
	})

	reminderGroup := e.Group("/reminder")
	reminderGroup.Use(middleware.JWT(signingKey))
	reminderGroup.POST("", handlers.CreateReminder)

	e.GET("/health-check", handlers.HealthCheck)

	v1 := e.Group("/v1")
	v1.POST("/login", handlers.Login)
	v1.POST("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}
