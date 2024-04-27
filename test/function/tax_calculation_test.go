package function

import (
	"github.com/son1122/assessment-tax/constant"
	"github.com/son1122/assessment-tax/db"
	"testing"

	"github.com/son1122/assessment-tax/util"
	"github.com/stretchr/testify/assert"
)

func TestTaxCalculation(t *testing.T) {
	constant.InitConfig()
	cfg := constant.Get()
	db.InitDB(cfg.DatabaseURL)
	// Define test cases
	testCases := []struct {
		name     string
		input    float64
		expected float64
	}{
		{name: "Income is exactly the threshold", input: 150000, expected: 0},
		{name: "Income just above threshold", input: 150001, expected: 0.1},
		{name: "Medium income", input: 500000, expected: 35000},
		{name: "Medium income just above bracket", input: 500001, expected: 35000.15},
		{name: "High income", input: 1000000, expected: 110000},
		{name: "High income just above bracket", input: 1000001, expected: 110000.2},
		{name: "Very high income", input: 2000000, expected: 310000},
		{name: "Very high income just above", input: 2000001, expected: 310000.35},
		{name: "Negative income", input: -1, expected: 0},
		{name: "Zero income", input: 0, expected: 0},
		{name: "Extremely high income", input: 10000000, expected: 3110000},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, actualResponse := util.TaxCalculationFromTotalIncome(tc.input)
			assert.Equal(t, tc.expected, actualResponse, "Unexpected tax calculation for input")
		})
	}
}
