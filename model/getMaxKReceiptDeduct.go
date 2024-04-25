package model

import (
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"
)

// GetPersonalDeduct query personal deduct from database
func GetMaxKReceiptDeduct() (float64, error) {

	var personalDeduct structs.GetTaxDeductStruct
	err := db.DB.QueryRow(`SELECT amount_deduct,id,is_active,create_at FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'max-k-receipt'`).Scan(&personalDeduct.PersonalDeduct, &personalDeduct.Id, &personalDeduct.Is_active, &personalDeduct.Create_at)
	if err != nil {
		return 0, err
	}
	return personalDeduct.PersonalDeduct, err
}
