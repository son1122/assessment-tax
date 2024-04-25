package model

import (
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"
)

// GetPersonalDeduct query personal deduct from database
func SetDonationDeduct() (float64, error) {

	var donationDeduct structs.GetTaxDeductStruct
	err := db.DB.QueryRow(`SELECT amount_deduct,id,is_active,create_at,version FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'donation'`).Scan(&donationDeduct.PersonalDeduct, &donationDeduct.Id, &donationDeduct.Is_active, &donationDeduct.Create_at, &donationDeduct.Version)
	if err != nil {
		return 0, err
	}
	return donationDeduct.PersonalDeduct, err
}
