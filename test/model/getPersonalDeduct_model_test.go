package model_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPersonalDeduct(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	expectedDeduct := 5000.00
	rows := sqlmock.NewRows([]string{"personal_deduct", "id", "is_active", "create_at"}).
		AddRow(expectedDeduct, 1, true, "2022-01-01")

	mock.ExpectQuery(`SELECT amount_deduct,id,is_active,create_at FROM public."master_deduct" WHERE is_active = TRUE AND type_deduct = 'personal'`).
		WillReturnRows(rows)
	db.SetDB(dbMock)

	personalDeduct, err := model.GetPersonalDeduct()

	assert.NoError(t, err)
	assert.Equal(t, expectedDeduct, personalDeduct)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	db.SetDB(nil)
}
