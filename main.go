package main

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/router"
	echoswagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
)

func main() {
	//e := echo.New()
	//e.GET("/", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	//})
	//e.Logger.Fatal(e.Start(":1323"))

	e := echo.New()

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

	// Use environment variable for the server port

	e.Logger.Fatal(e.Start(":" + PORT))
}
