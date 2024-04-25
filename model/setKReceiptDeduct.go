package model

import (
	"errors"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"
)

// GetPersonalDeduct query personal deduct from database
func SetKReceiptDeduct(kReceiptDeduct float64) (float64, error) {

	// Start a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return 0, err
	}

	// Select the current active record
	var donationDeduct structs.GetTaxDeductStruct
	err = tx.QueryRow(`SELECT amount_deduct, id, is_active, create_at, version FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'k-receipt'`).Scan(&donationDeduct.PersonalDeduct, &donationDeduct.Id, &donationDeduct.Is_active, &donationDeduct.Create_at, &donationDeduct.Version)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	maxKReceipt, _ := GetMaxKReceiptDeduct()
	floorKReceipt, _ := GetFloorKReceiptDeduct()
	if kReceiptDeduct > maxKReceipt || kReceiptDeduct < floorKReceipt {
		return donationDeduct.PersonalDeduct, errors.New("donation is not in maximum or minimum range")
	}

	// Set the current active record to inactive
	_, err = tx.Exec(`UPDATE public."master_deduct" SET is_active = FALSE WHERE id = $1`, donationDeduct.Id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Insert a new record with an incremented version and is_active = TRUE
	newVersion := donationDeduct.Version + 1
	_, err = tx.Exec(`INSERT INTO public."master_deduct" (amount_deduct, is_active, type_deduct, version) VALUES ($1, $2 , 'k-receipt', $3)`, kReceiptDeduct, true, newVersion)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return kReceiptDeduct, nil
}
