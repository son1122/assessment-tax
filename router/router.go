package router

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
)

func InitRoutes(e *echo.Echo) {

	e.POST("/tax/calculations", controller.TaxCalculationPost)
	e.POST("/admin/deductions/personal", controller.AdminDeductionPersonalAdjust)
	e.POST("/tax/calculations/upload-csv", controller.TaxCalculationCSVPost)

}
