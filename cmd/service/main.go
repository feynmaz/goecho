package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/feynmaz/goecho/handlers"
	"github.com/feynmaz/goecho/models"
)

func main() {
	e := echo.New()

	// Signing Key for the auth middleware
	var signingKey = []byte("secret")
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

	e.GET("/healthcheck", handlers.HealthCheck)

	g := e.Group("/v1")
	g.POST("/login", handlers.Login)
	g.POST("/logout", handlers.Logout)

	e.Logger.Fatal(e.Start(":8080"))
}
