package structs

type AdminResponseStruct struct {
	PersonalDeduction float64 `json:"personalDeduction,omitempty" validate:"gte=0"`
	KReceipt          float64 `json:"kReceipt,omitempty" validate:"gte=0"`
	Donation          float64 `json:"donation,omitempty" validate:"gte=0"`
}
