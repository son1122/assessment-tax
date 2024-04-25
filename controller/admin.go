package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/model"
	struc "github.com/son1122/assessment-tax/structs"
	"net/http"
)

// AdminDeductionPersonalAdjust handles the POST /admin/deductions/personal route
// @Summary AdminDeductionPersonalAdjust
// @Description AdminDeductionPersonalAdjust
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   tax_body  body      _struct.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} _struct.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]

func AdminDeductionPersonalAdjust(c echo.Context) error {
	var amount struc.AdminRequestStruct
	if err := c.Bind(&amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if err := c.Validate(amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	setDonationResult, err := model.SetDonationDeduct(amount.Amount)

	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	response := struc.AdminResponseStruct{
		Donation: setDonationResult,
	}
	return c.JSON(http.StatusOK, response)

}

// AdminDeductionPersonalAdjust handles the POST /admin/deductions/personal route
// @Summary AdminDeductionPersonalAdjust
// @Description AdminDeductionPersonalAdjust
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   tax_body  body      _struct.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} _struct.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]

func AdminDeductionKReceiptAdjust(c echo.Context) error {
	var amount struc.AdminRequestStruct
	if err := c.Bind(&amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if err := c.Validate(amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	setDonationResult, err := model.SetDonationDeduct(amount.Amount)

	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	response := struc.AdminResponseStruct{
		Donation: setDonationResult,
	}
	return c.JSON(http.StatusOK, response)

}

// AdminDeductionPersonalAdjust handles the POST /admin/deductions/personal route
// @Summary AdminDeductionPersonalAdjust
// @Description AdminDeductionPersonalAdjust
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   tax_body  body      _struct.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} _struct.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]

func AdminDeductionDonationAdjust(c echo.Context) error {
	var amount struc.AdminRequestStruct
	if err := c.Bind(&amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	if err := c.Validate(amount); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	setDonationResult, err := model.SetDonationDeduct(amount.Amount)

	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	response := struc.AdminResponseStruct{
		Donation: setDonationResult,
	}
	return c.JSON(http.StatusOK, response)

}
