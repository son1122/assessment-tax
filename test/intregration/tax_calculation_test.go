package integration

import (
	"bytes"
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

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{
    "taxRefund": 486000,
    "taxLevel": [
        {
            "level": "0 - 150000",
            "tax": 0
        },
        {
            "level": "150000 - 500000",
            "tax": 14000
        },
        {
            "level": "500000 - 1000000",
            "tax": 0
        },
        {
            "level": "1000000 - 2000000",
            "tax": 0
        },
        {
            "level": "2000000 - ขึ้นไป",
            "tax": 0
        }
    ]
}`
	if !bytes.Contains(body, []byte(expected)) {
		t.Errorf("Unexpected response body: got %v want %v", string(body), expected)
	}
}
