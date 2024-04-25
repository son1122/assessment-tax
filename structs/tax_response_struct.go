package structs

type TaxResponse struct {
	Tax       float64        `json:"tax,omitempty"`
	TaxRefund float64        `json:"taxRefund,omitempty"`
	TaxLevel  []TaxLevelData `json:"taxLevel,omitempty"`
}
