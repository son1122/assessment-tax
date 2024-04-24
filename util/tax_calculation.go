package util

import (
	"github.com/son1122/assessment-tax/model"
	"github.com/son1122/assessment-tax/structs"
)

// TaxCalculationFromTotalIncome Function for calculate tax from total income by using Tax Level From Database from model.GetTaxLevel()
func TaxCalculationFromTotalIncome(income float64) ([]structs.TaxLevelData, float64) {

	taxLevel, _ := model.GetTaxLevel()
	var tax float64 = 0
	var taxValueInLevel float64 = 0
	//var sum float64 = 0
	for i := 0; i < len(taxLevel); i++ {
		if i == len(taxLevel)-1 {
			taxValueInLevel = income - float64(taxLevel[i].Floor)
			tax = tax + taxValueInLevel*float64(taxLevel[i].TaxValue)/100
			break
		}
		if income >= float64(taxLevel[i].Ceil) {
			if i > 0 {
				tax = tax + (float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil) - float64(taxLevel[i-1].Ceil)) / 100)
			} else {
				tax = tax + (float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil)) / 100)
			}

		} else {
			if income <= float64(taxLevel[i].Floor) {
				break
			}
			taxValueInLevel = income - float64(taxLevel[i].Floor)
			tax = tax + taxValueInLevel*float64(taxLevel[i].TaxValue)/100
			break
		}

	}
	//log.Printf(string(tax))
	data := []structs.TaxLevelData{
		{Level: "0-150,000", Tax: 0},
		{Level: "0-150,001", Tax: 500},
	}
	return data, tax

}
