package structs

type TaxResponseCSVDataStruct struct {
	TotalIncome float64 `json:"totalIncome" validate:"gte=0"`
	Tax         float64 `json:"tax" validate:"gte=0"`
}

type TaxResponseCSVStruct struct {
	Taxes []TaxResponseCSVDataStruct `json:"taxes" validate:"gte=0"`
}
