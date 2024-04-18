package _struct

type TaxStruct struct {
	TotalIncome float64           `json:"totalIncome"`
	Wht         float64           `json:"wht"`
	Allowances  []AllowanceStruct `json:"allowances"`
}
type AllowanceStruct struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}
