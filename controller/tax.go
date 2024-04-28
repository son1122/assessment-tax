package controller

import (
	"encoding/csv"
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/model"
	struc "github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"io"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"
)

// TaxCalculationPost handles the POST /tax/calculations route
// @Summary Calculate taxes
// @Description Calculates taxes based on total income, withholding tax, and allowances.
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   tax_body  body      struc.TaxStruct  true  "Tax Calculation Request"
// @Success 200 {object} struc.TaxResponse  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations [post]
func TaxCalculationPost(c echo.Context) error {
	log.Println("start TaxCalculationPost")
	var tax struc.TaxStruct
	err := c.Bind(&tax)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}
	err = c.Validate(tax)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	deductTypeAndAmount, err := model.GetDeductType()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	totalAllowance := 0.0

	for _, typeDeduct := range tax.Allowances {
		deductTypeAndAmount[typeDeduct.AllowanceType] += typeDeduct.Amount
		totalAllowance += typeDeduct.Amount
	}

	personalDeduct, err := model.GetPersonalDeduct()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	donationDeduct, err := model.GetDonationDeduct()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	kReceiptDeduct, err := model.GetKReceiptDeduct()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	if donationDeduct < deductTypeAndAmount["donation"] {
		deductTypeAndAmount["donation"] = donationDeduct
	}
	if kReceiptDeduct < deductTypeAndAmount["k-receipt"] {
		deductTypeAndAmount["k-receipt"] = kReceiptDeduct
	}
	incomeDeductPersonal := tax.TotalIncome - personalDeduct
	incomeDeductDonation := incomeDeductPersonal - deductTypeAndAmount["donation"]
	incomeDeductKreceipt := incomeDeductDonation - deductTypeAndAmount["k-receipt"]
	taxLevelData, taxCost := util.TaxCalculationFromTotalIncome(incomeDeductKreceipt)
	finalTax := taxCost - tax.Wht
	log.Println("fin TaxCalculationPost", finalTax)
	if finalTax >= 0 {
		taxResponse := struc.TaxResponse{Tax: strconv.FormatFloat(finalTax, 'f', 2, 64), TaxLevel: taxLevelData}
		return c.JSON(http.StatusOK, taxResponse)
	} else {
		taxResponse := struc.TaxResponse{TaxRefund: math.Abs(finalTax), TaxLevel: taxLevelData}
		return c.JSON(http.StatusOK, taxResponse)
	}

}

// TaxCalculationCSVPost handles the POST /tax/calculations/upload-csv route
// @Summary Calculate taxes CSV
// @Description Calculates taxes based on total income, withholding tax, and allowances. CSV
// @Tags tax
// @Accept  json
// @Produce  json
// @Param   tax_body  body      FormFile  true  "Tax Calculation Request"
// @Success 200 {object} struc.TaxResponseCSVStruct  "Returns the calculated tax amount"
// @Failure 400 {string} string "Invalid input parameters"
// @Router /tax/calculations/upload-csv [post]
func TaxCalculationCSVPost(c echo.Context) error {
	log.Println("TaxCalculationCSVPost")
	file, err := c.FormFile("taxFile")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get the file")
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) error {
		err := src.Close()
		if err != nil {
			return err
		}
		return nil
	}(src)

	reader := csv.NewReader(src)
	var taxes []struc.TaxResponseCSVDataStruct

	if _, err = reader.Read(); err != nil {
		return err
	}
	log.Println("file read success")
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
		if finalTax >= 0 {
			taxes = append(
				taxes,
				struc.TaxResponseCSVDataStruct{
					TotalIncome: totalIncome,
					Tax:         finalTax,
				},
			)
		} else {
			taxes = append(
				taxes,
				struc.TaxResponseCSVDataStruct{
					TotalIncome: totalIncome,
					TaxRefund:   math.Abs(finalTax),
				},
			)
		}

		loopNumber += 1
	}
	log.Println("Fin TaxCalculationCSVPost")
	return c.JSON(http.StatusOK, struc.TaxResponseCSVStruct{Taxes: taxes})
}
