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
	setPersonalResult, err := model.SetPersonalDeduct(amount.Amount)

	if err != nil && setPersonalResult != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if setPersonalResult == 0 {
		return c.JSON(http.StatusOK, map[string]int{"personalDeduction": 0})
	}
	response := struc.AdminResponseStruct{
		PersonalDeduction: setPersonalResult,
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

	setKReceiptResult, err := model.SetKReceiptDeduct(amount.Amount)

	if err != nil && setKReceiptResult != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if setKReceiptResult == 0 {
		return c.JSON(http.StatusOK, map[string]int{"kReceipt": 0})
	}
	response := struc.AdminResponseStruct{
		KReceipt: setKReceiptResult,
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

	if err != nil && setDonationResult != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if setDonationResult == 0 {
		return c.JSON(http.StatusOK, map[string]int{"donationDeduction": 0})
	}
	response := struc.AdminResponseStruct{
		Donation: setDonationResult,
	}
	return c.JSON(http.StatusOK, response)

}
