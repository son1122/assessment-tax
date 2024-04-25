package router

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
	"github.com/son1122/assessment-tax/middleware"
)

func InitRoutes(e *echo.Echo) {

	//TaxCalculation Controller
	e.POST("/tax/calculations", controller.TaxCalculationPost)
	e.POST("/tax/calculations/upload-csv", controller.TaxCalculationCSVPost)

	//AdminCalculation Controller
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.BasicAuthMiddleware)
	adminGroup.POST("/deductions/k-receipt", controller.AdminDeductionKReceiptAdjust)
	adminGroup.POST("/deductions/donation", controller.AdminDeductionDonationAdjust)
	adminGroup.POST("/deductions/personal", controller.AdminDeductionPersonalAdjust)
}
