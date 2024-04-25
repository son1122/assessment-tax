package structs

type AdminRequestStruct struct {
	Amount float64 `json:"amount" validate:"gte=0"`
}
