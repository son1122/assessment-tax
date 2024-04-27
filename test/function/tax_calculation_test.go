package function

import (
	"encoding/json"
	"github.com/son1122/assessment-tax/db"
	"github.com/son1122/assessment-tax/structs"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCalculation(t *testing.T) {
	DatabaseUrl := os.Getenv("DatabaseUrl")
	if DatabaseUrl == "" {
		DatabaseUrl = "host=localhost port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"
		db.InitDB(DatabaseUrl)
	}
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Scenario 1: Total income less than or equal to 150000",
			input: `{
				"totalIncome": 150000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 0.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 0.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 2: Total income between 150001 and 500000 ",
			input: `{
				"totalIncome": 500000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 29000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 29000.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 3: Total income between 500001 and 1000000",
			input: `{
				"totalIncome": 1000000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 101000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 35000.0},
					{"level": "500,001-1,000,000", "tax": 66000.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 4: Total income between 1000000 and 2000000",
			input: `{
				"totalIncome": 2000000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 298000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 35000.0},
					{"level": "500,001-1,000,000", "tax": 75000.0},
					{"level": "1,000,001-2,000,000", "tax": 188000.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 5: Total income more than 2000001",
			input: `{
				"totalIncome": 2060005.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 310001.75,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 35000.0},
					{"level": "500,001-1,000,000", "tax": 75000.0},
					{"level": "1,000,001-2,000,000", "tax": 200000.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 1.75}
				]
			}`,
		},
		{
			name: "Scenario 6: Total income between 150001 and 500000 with wht",
			input: `{
				"totalIncome": 500000.0,
				"wht": 30000.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 0.0,
				"taxRefund ": 1000,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 29000.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 7: Total income between 150001 and 500000 with allowance more than limit",
			input: `{
				"totalIncome": 500000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 100000.0
					},
					{
						"allowanceType": "donation",
						"amount": 200000.0
					}
				]
			}`,
			expected: `{
				"tax": 14000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 14000.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 8: Total income between 150001 and 500000 with wht and allowance more than limit ",
			input: `{
				"totalIncome": 500000.0,
				"wht": 20000.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 100000.0
					},
					{
						"allowanceType": "donation",
						"amount": 200000.0
					}
				]
			}`,
			expected: `{
				"tax": 0.0,
				"taxRefund ": 6000,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 14000.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var requestBody structs.TaxResponse
			if err := json.Unmarshal([]byte(tc.input), &requestBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}

			expectedResponse := structs.TaxResponse{}
			if err := json.Unmarshal([]byte(tc.expected), &expectedResponse); err != nil {
				t.Fatalf("failed to unmarshal expected response: %v", err)
			}
			allowance, _, _ := util.AllowanceCalculation(requestBody)
			requestBody.TotalIncome -= allowance
			actualResponse := TaxCalculation(requestBody.TotalIncome, requestBody.Wht)

			assert.Equal(t, expectedResponse, actualResponse, "Response does not match expected")
		})
	}
}
