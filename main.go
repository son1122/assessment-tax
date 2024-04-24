package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/son1122/assessment-tax/db"
	_ "github.com/son1122/assessment-tax/docs"
	"github.com/son1122/assessment-tax/router"
	"github.com/son1122/assessment-tax/util"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Use environment variables for database connection
	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable" // Default port if not specified
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port if not specified
	}
	ADMIN_USERNAME := os.Getenv("ADMIN_USERNAME")
	if ADMIN_USERNAME == "" {
		ADMIN_USERNAME = "adminTax" // Default port if not specified
	}
	ADMIN_PASSWORD := os.Getenv("ADMIN_PASSWORD")
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "admin!" // Default port if not specified
	}

	db.InitDB(DATABASE_URL)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db.DB)
			return next(c)
		}
	})

	router.InitRoutes(e)

	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	go func() {
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("error shutting down the server: ", err)
	}
	e.Logger.Info("Server gracefully stopped")
	fmt.Println("shutting down the server")

}
