package _struct

type TaxStruct struct {
	TotalIncome float64           `json:"totalIncome" validate:"gte=0"`
	Wht         float64           `json:"wht" validate:"gte=0"`
	Allowances  []AllowanceStruct `json:"allowances" validate:"dive"`
}
type AllowanceStruct struct {
	AllowanceType string  `json:"allowanceType" validate:"oneof=donation k-receipt"`
	Amount        float64 `json:"amount" validate:"gte=0"`
}
