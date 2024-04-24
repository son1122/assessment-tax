package _test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/son1122/assessment-tax/controller"
	"github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaxCalculationPost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setting up the SQL Mock
	expectedDeduct := 5000.00
	rows := sqlmock.NewRows([]string{"personal_deduct"}).AddRow(expectedDeduct)
	mock.ExpectQuery(`SELECT personal_deduct FROM public."master_personal_deduct" WHERE is_active = TRUE`).WillReturnRows(rows)

	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	reqTax := structs.TaxStruct{
		TotalIncome: 100000,
		Wht:         1000,
		Allowances: []structs.AllowanceStruct{
			{AllowanceType: "donation", Amount: 500},
		},
	}

	reqBody, _ := json.Marshal(reqTax)
	req := httptest.NewRequest(http.MethodPost, "/tax/calculations", bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("db", db) // Pass the mock DB to the context
	c.SetPath("/tax/calculations")

	// Invoke the handler
	if assert.NoError(t, controller.TaxCalculationPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var response structs.TaxResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
			t.Fatal("Failed to unmarshal response:", err)
		}

		expectedTax, _ := util.TaxCalculationFromTotalIncome(reqTax.TotalIncome - expectedDeduct) // Adjusting income by the mocked deduction
		assert.Equal(t, expectedTax, response.Tax)
	}

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
