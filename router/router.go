package router

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
)

func InitRoutes(e *echo.Echo) {
	//user
	e.GET("/users", controller.GetUsers)
	e.POST("/tax/calculations", controller.TaxCalculation)

}
