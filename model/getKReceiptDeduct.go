package model

import (
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"
)

// GetPersonalDeduct query personal deduct from database
func GetKReceiptDeduct() (float64, error) {

	var kReceiptDeduct structs.GetTaxDeductStruct
	err := db.DB.QueryRow(`SELECT amount_deduct,id,is_active,create_at FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'k-receipt'`).Scan(&kReceiptDeduct.PersonalDeduct, &kReceiptDeduct.Id, &kReceiptDeduct.Is_active, &kReceiptDeduct.Create_at)
	if err != nil {
		return 0, err
	}
	return kReceiptDeduct.PersonalDeduct, err
}
