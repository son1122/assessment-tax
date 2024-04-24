package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/model"
	struc "github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"net/http"
)

// TaxCalculationPost handles the POST /tax/calculations route
// @Summary Calculate taxes
// @Description Calculates taxes based on total income, withholding tax, and allowances.
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   tax_body  body      _struct.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} _struct.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]
func TaxCalculationPost(c echo.Context) error {
	var tax struc.TaxStruct
	if err := c.Bind(&tax); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if err := c.Validate(tax); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	personalDeduct, _ := model.GetPersonalDeduct()
	incomeDeductPersonal := tax.TotalIncome - personalDeduct
	_, taxCost := util.TaxCalculationFromTotalIncome(incomeDeductPersonal)
	taxResponse := struc.TaxResponse{Tax: taxCost}

	return c.JSON(http.StatusOK, taxResponse)
}
