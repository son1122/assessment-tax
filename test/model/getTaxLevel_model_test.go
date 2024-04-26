package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/model"
	"testing"
)

func TestGetTaxLevel(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	query := "SELECT id, floor, ceil, create_at,tax_value FROM public.\"master_tax_level\" ORDER BY floor"

	// Successful query setup
	rows := sqlmock.NewRows([]string{"id", "floor", "ceil", "create_at", "tax_value"}).
		AddRow(1, 0, 150000, "2022-01-01", 5).
		AddRow(2, 150001, 300000, "2022-01-01", 10)
	mock.ExpectQuery(query).WillReturnRows(rows)
	db.SetDB(dbMock)
	taxLevels, err := model.GetTaxLevel()
	if err != nil {
		t.Error("Failed to get tax levels:", err)
	}
	if len(taxLevels) != 2 {
		t.Errorf("Expected 2 tax levels, got %d", len(taxLevels))
	}

	// Test query failure
	mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
	_, err = model.GetTaxLevel()
	if err == nil {
		t.Error("Expected an error when query fails, but got nil")
	}

	// Test row scan failure due to incorrect type
	brokenRows := sqlmock.NewRows([]string{"id", "floor", "ceil", "create_at", "tax_value"}).
		AddRow("wrong_type", "not_int", "not_int", "not_date", "not_float").
		AddRow("wrong_type", "not_int", "not_int", "not_date", "not_float").
		AddRow("wrong_type", "not_int", "not_int", "not_date", "not_float").
		AddRow("wrong_type", "not_int", "not_int", "not_date", "not_float").
		AddRow("wrong_type", "not_int", "not_int", "not_date", "not_float")

	mock.ExpectQuery(query).WillReturnRows(brokenRows)

	_, err = model.GetTaxLevel()
	if err == nil {
		t.Error("Expected an error during row scan, but got nil")
	}

	// Check that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
	db.SetDB(nil)
}
