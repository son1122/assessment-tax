package util

import (
	"fmt"
	"github.com/son1122/assessment-tax/model"
	"github.com/son1122/assessment-tax/structs"
)

// TaxCalculationFromTotalIncome Function for calculate tax from total income by using Tax Level From Database from model.GetTaxLevel()
func TaxCalculationFromTotalIncome(income float64) ([]structs.TaxLevelData, float64) {

	data := []structs.TaxLevelData{}

	taxLevel, _ := model.GetTaxLevel()
	var tax float64 = 0
	var incomeValueInLevel float64 = 0
	//var sum float64 = 0
	for i := 0; i < len(taxLevel); i++ {
		if i == len(taxLevel)-1 {
			if income >= float64(taxLevel[i].Floor) {
				incomeValueInLevel = income - float64(taxLevel[i].Floor)
				taxValueInLevel := incomeValueInLevel * float64(taxLevel[i].TaxValue) / 100
				tax = tax + taxValueInLevel
				data = append(data, structs.TaxLevelData{
					Level: fmt.Sprintf("%s ขึ้นไป", commaFormat(taxLevel[i].Floor+1)),
					Tax:   taxValueInLevel,
				})
				break
			}
			data = append(data, structs.TaxLevelData{
				Level: fmt.Sprintf("%s ขึ้นไป", commaFormat(taxLevel[i].Floor+1)),
				Tax:   0,
			})
			break
		}
		if income >= float64(taxLevel[i].Ceil) {
			if i > 0 {
				taxValueInLevel := float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil) - float64(taxLevel[i-1].Ceil)) / 100
				tax = tax + taxValueInLevel
				data = append(data, structs.TaxLevelData{
					Level: fmt.Sprintf("%s-%s", commaFormat(taxLevel[i].Floor+1), commaFormat(taxLevel[i].Ceil)),
					Tax:   taxValueInLevel,
				})
			} else {
				taxValueInLevel := float64(taxLevel[i].TaxValue) * (float64(taxLevel[i].Ceil)) / 100
				tax = tax + taxValueInLevel
				data = append(data, structs.TaxLevelData{
					Level: fmt.Sprintf("%s-%s", commaFormat(taxLevel[i].Floor), commaFormat(taxLevel[i].Ceil)),
					Tax:   taxValueInLevel,
				})
			}

		} else {
			if income <= float64(taxLevel[i].Floor) {
				data = append(data, structs.TaxLevelData{
					Level: fmt.Sprintf("%s-%s", commaFormat(taxLevel[i].Floor+1), commaFormat(taxLevel[i].Ceil)),
					Tax:   0,
				})
				continue
			}

			incomeValueInLevel = income - float64(taxLevel[i].Floor)
			taxValueInLevel := incomeValueInLevel * float64(taxLevel[i].TaxValue) / 100
			tax = tax + taxValueInLevel
			data = append(data, structs.TaxLevelData{
				Level: fmt.Sprintf("%s-%s", commaFormat(taxLevel[i].Floor+1), commaFormat(taxLevel[i].Ceil)),
				Tax:   taxValueInLevel,
			})
			continue
		}

	}

	return data, tax

}
