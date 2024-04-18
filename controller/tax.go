package controller

import (
	"github.com/labstack/echo/v4"
	model "github.com/son1122/assessment-tax/model"
	"net/http"
)

type taxLevelData struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

func TaxCalculationFromTotalIncome(totalIncome float64) ([]taxLevelData, float64) {

	totalIncome = 5000000
	taxLevel, _ := model.GetTaxLevel()
	var tax float64 = 0
	var taxValueInLevel float64 = 0
	//var sum float64 = 0
	for i := 0; i < len(taxLevel); i++ {
		if i == len(taxLevel)-1 {
			taxValueInLevel = totalIncome - float64(taxLevel[i].Floor)
			tax = tax + taxValueInLevel*float64(taxLevel[i].TaxValue)/100
			break
		}
		if totalIncome >= float64(taxLevel[i].Ceil) {
			if i > 0 {
				tax = tax + (float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil) - float64(taxLevel[i-1].Ceil)) / 100)
			} else {
				tax = tax + (float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil)) / 100)
			}

		} else {
			if totalIncome <= float64(taxLevel[i].Floor) {
				break
			}
			taxValueInLevel = totalIncome - float64(taxLevel[i].Floor)
			tax = tax + taxValueInLevel*float64(taxLevel[i].TaxValue)/100
			break
		}

	}
	//log.Printf(string(tax))
	data := []taxLevelData{
		taxLevelData{Level: "0-150,000", Tax: 0},
		taxLevelData{Level: "0-150,001", Tax: 500},
	}
	return data, tax

}

// TaxCalculationPost handles the POST tax/calculations route
// @Summary tax/calculations
// @Description tax/calculations by data
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /tax/calculations [post]
func TaxCalculationPost(c echo.Context) error {
	_, tax := TaxCalculationFromTotalIncome(123.55)
	//data, _ := model.GetPersonalDeduct()
	//data, _ := model.GetTaxLevel()
	return c.JSON(http.StatusOK, tax)
}
