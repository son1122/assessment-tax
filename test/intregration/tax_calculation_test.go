package integration

import (
	"encoding/json"
	"fmt"
	struc "github.com/son1122/assessment-tax/structs"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// TestTaxCalculationFunction tests the POST request handling of the tax calculation API.
func TestTaxCalculationFunction(t *testing.T) {

	url := `http://localhost:8080/tax/calculations`
	method := "POST"

	payload := strings.NewReader(`{
        "totalIncome": 500000.0,
        "wht": 500000,
        "allowances": [
            {
                "allowanceType": "k-receipt",
                "amount": 100000.0
            },
            {
                "allowanceType": "donation",
                "amount": 100000.0
            }
        ]
    }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var actualResponse struc.TaxResponse
	if err := json.Unmarshal(body, &actualResponse); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	expectedResponse := struc.TaxResponse{
		TaxRefund: 486000,
		TaxLevel: []struc.TaxLevelData{
			{Level: "0-150,000", Tax: 0},
			{Level: "150,001-500,000", Tax: 14000},
			{Level: "500,001-1,000,000", Tax: 0},
			{Level: "1,000,001-2,000,000", Tax: 0},
			{Level: "2,000,001 ขึ้นไป", Tax: 0},
		},
	}
	if !compareTaxResponses(expectedResponse, actualResponse) {
		t.Errorf("Unexpected response: got %+v want %+v", actualResponse, expectedResponse)
	}

}

func compareTaxResponses(a, b struc.TaxResponse) bool {
	if a.Tax != b.Tax || a.TaxRefund != b.TaxRefund || len(a.TaxLevel) != len(b.TaxLevel) {
		return false
	}
	for i := range a.TaxLevel {
		if a.TaxLevel[i].Level != b.TaxLevel[i].Level || a.TaxLevel[i].Tax != b.TaxLevel[i].Tax {
			fmt.Printf("Mismatch found in TaxLevel[%d]: Expected %v, got %v\n", i, a.TaxLevel[i], b.TaxLevel[i])
			return false
		}
	}
	return true
}
