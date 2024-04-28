package structs

type TaxResponse struct {
	Tax       string         `json:"tax,omitempty"`
	TaxRefund float64        `json:"taxRefund,omitempty"`
	TaxLevel  []TaxLevelData `json:"taxLevel,omitempty"`
}
