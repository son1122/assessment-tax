package model

import (
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"
)

// GetTaxLevel query tax level from database
func GetTaxLevel() ([]structs.TaxLevel, error) {

	rows, err := db.DB.Query(`SELECT id, floor, ceil, create_at,tax_value FROM public."master_tax_level" ORDER BY floor`)
	if err != nil {
		return nil, err // Return the error to be handled by the caller
	}
	defer rows.Close()

	var taxLevels []structs.TaxLevel
	for rows.Next() {
		var tl structs.TaxLevel
		if err := rows.Scan(&tl.ID, &tl.Floor, &tl.Ceil, &tl.CreateAt, &tl.TaxValue); err != nil {
			return nil, err // Return immediately on error
		}
		taxLevels = append(taxLevels, tl)
	}

	if err = rows.Err(); err != nil {
		return nil, err // Handle any errors encountered during iteration
	}

	return taxLevels, nil
}
