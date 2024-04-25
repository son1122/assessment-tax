package controller

import (
	"encoding/csv"
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/model"
	struc "github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"io"
	"net/http"
	"strconv"
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
	totalDonationAllowance := 0.0
	for _, allowance := range tax.Allowances {
		totalDonationAllowance += allowance.Amount
	}

	personalDeduct, _ := model.GetPersonalDeduct()
	donationDeduct, _ := model.GetDonationDeduct()
	incomeDeductPersonal := tax.TotalIncome - personalDeduct

	if totalDonationAllowance > donationDeduct {
		totalDonationAllowance = donationDeduct
	}
	incomeDeductDonation := incomeDeductPersonal - totalDonationAllowance
	taxLevelData, taxCost := util.TaxCalculationFromTotalIncome(incomeDeductDonation)
	finalTax := taxCost - tax.Wht
	if finalTax >= 0 {
		taxResponse := struc.TaxResponse{Tax: finalTax, TaxLevel: taxLevelData}
		return c.JSON(http.StatusOK, taxResponse)
	} else {
		taxResponse := struc.TaxResponse{TaxRefund: finalTax, TaxLevel: taxLevelData}
		return c.JSON(http.StatusOK, taxResponse)
	}

}

// TaxCalculationCSVPost handles the POST /tax/calculations route
// @Summary Calculate taxes CSV
// @Description Calculates taxes based on total income, withholding tax, and allowances. CSV
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   tax_body  body      _struct.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} _struct.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]
func TaxCalculationCSVPost(c echo.Context) error {

	file, err := c.FormFile("taxFile")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get the file")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	reader := csv.NewReader(src)
	var taxes []struc.TaxResponseCSVDataStruct

	if _, err = reader.Read(); err != nil {
		return err
	}
	var loopNumber = 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read the csv file")
		}

		totalIncome, err := strconv.ParseFloat(record[0], 64)
		wht, err := strconv.ParseFloat(record[1], 64)
		donation, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid totalIncome value")
		}
		personalDeduct, _ := model.GetPersonalDeduct()
		//donationDeduct, _ := model.GetDonationDeduct()
		totalIncomeDeductPersonal := totalIncome - personalDeduct
		totalIncomeDeductDonation := totalIncomeDeductPersonal - donation
		_, tax := util.TaxCalculationFromTotalIncome(totalIncomeDeductDonation)
		finalTax := tax - wht
		taxes = append(
			taxes,
			struc.TaxResponseCSVDataStruct{
				TotalIncome: totalIncome,
				Tax:         finalTax,
			},
		)

		loopNumber += 1
	}

	return c.JSON(http.StatusOK, struc.TaxResponseCSVStruct{Taxes: taxes})
}
