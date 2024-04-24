package model

import (
	"github.com/son1122/assessment-tax/structs"
	"github.com/son1122/assessment-tax/util"
	"reflect"
	"testing"
)

func TestTaxCalculationFromTotalIncome(t *testing.T) {
	type args struct {
		ncome float64
	}
	tests := []struct {
		name  string
		args  args
		want  []structs.TaxLevelData
		want1 float64
	}{
		// TODO: Add test cases.
		{"test", args{ncome: 1000000}, nil, 15000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := util.TaxCalculationFromTotalIncome(tt.args.ncome)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaxCalculationFromTotalIncome() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TaxCalculationFromTotalIncome() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
