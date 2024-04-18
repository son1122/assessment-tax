package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/router"
	util "github.com/son1122/assessment-tax/util"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"os/signal"
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

	router.InitRoutes(e)

	e.GET("/swagger/*", echoswagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	go func() {
		if err := e.Start(":" + PORT); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
