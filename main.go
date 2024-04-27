package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/son1122/assessment-tax/constant"
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
	e.Logger.SetLevel(log.DEBUG)

	constant.InitConfig()
	cfg := constant.Get()

	db.InitDB(cfg.DatabaseURL)
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy")
	})
	e.GET("/swagger/*", echoswagger.WrapHandler)
	router.InitRoutes(e)

	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	time.Sleep(5 * time.Second) // Introduces a 5-second delay after Ctrl+C is pressed.

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("error shutting down the server: ", err)
	}

	fmt.Println("Shutting down the server")
}
