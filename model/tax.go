package model

import (
	"github.com/son1122/assessment-tax/db"
	"log"
)

func GetDonationDeduct() (float64, error) {

	var donationDeduct GetDonationDeductStruct
	err := db.DB.QueryRow(`SELECT donation_deduct,donation_id,is_active,create_at FROM public."master_donation_deduct" WHERE is_active = TRUE`).Scan(&donationDeduct.DonationDeduct, &donationDeduct.id, &donationDeduct.is_active, &donationDeduct.create_at)

	return donationDeduct.DonationDeduct, err
}

// GetPersonalDeduct query personal deduct from database
func GetPersonalDeduct() (float64, error) {

	var personalDeduct GetPersonalDeductStruct
	err := db.DB.QueryRow(`SELECT personal_deduct,id,is_active,create_at FROM public."master_personal_deduct" WHERE is_active = TRUE`).Scan(&personalDeduct.personalDeduct, &personalDeduct.id, &personalDeduct.is_active, &personalDeduct.create_at)

	return personalDeduct.personalDeduct, err
}

// GetTaxLevel query tax level from database
func GetTaxLevel() ([]TaxLevel, error) {

	rows, err := db.DB.Query(`SELECT id, floor, ceil, create_at,tax_value FROM public."master_tax_level" ORDER BY floor`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var taxLevels []TaxLevel
	for rows.Next() {
		var tl TaxLevel
		err := rows.Scan(&tl.ID, &tl.Floor, &tl.Ceil, &tl.CreateAt, &tl.TaxValue)
		if err != nil {
			log.Fatal(err)
		}
		taxLevels = append(taxLevels, tl)
	}
	return taxLevels, err
}
