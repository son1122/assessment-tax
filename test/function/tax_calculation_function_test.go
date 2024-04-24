package model_test

import (
	"github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"reflect"
	"testing"
)

func TestTaxCalculationFromTotalIncome(t *testing.T) {
	tests := []struct {
		name  string
		args  float64
		want  []structs.TaxLevelData
		want1 float64
	}{
		{
			name: "Test with income below first tax level",
			args: 100000, // Example total income
			want: []structs.TaxLevelData{
				{Level: "0-150,000", Tax: 0},
				{Level: "0-150,001", Tax: 500},
			}, // Expected tax levels
			want1: 0, // Expected tax amount
		},
		// Add more test cases for different income levels and expected tax amounts
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := util.TaxCalculationFromTotalIncome(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxCalculationFromTotalIncome() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TaxCalculationFromTotalIncome() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
