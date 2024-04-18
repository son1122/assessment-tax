package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetUsers handles the GET users route
// @Summary Show account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {string} string "ok"
// @Router /accounts/{id} [get]
func TaxCalculation(c echo.Context) error {

	return c.JSON(http.StatusOK, "tax calculation")
}
