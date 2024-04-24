package router

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
)

func InitRoutes(e *echo.Echo) {

	e.POST("/tax/calculations", controller.TaxCalculationPost)

}
