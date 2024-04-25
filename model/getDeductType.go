package model

import (
	"github.com/son1122/assessment-tax/db"
)

// GetDeductType query personal deduct from database
func GetDeductType() (map[string]float64, error) {
	deductAmount := make(map[string]float64)

	// Use Query to handle multiple rows

	rows, err := db.DB.Query(`
		SELECT DISTINCT ON (md.type_deduct) md.type_deduct
		FROM master_deduct md
		WHERE md.type_deduct NOT LIKE 'max-%'
		  AND md.type_deduct NOT LIKE 'floor-%'
		  AND md.is_active = 'true'
		ORDER BY md.type_deduct, md.create_at;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var typeDeduct string
	amount := 0.0

	for rows.Next() {
		if err := rows.Scan(&typeDeduct); err != nil {
			return nil, err
		}
		deductAmount[typeDeduct] = amount
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deductAmount, nil
}
