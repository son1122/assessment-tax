package model

type TaxLevel struct {
	ID       int
	Floor    int
	Ceil     int
	CreateAt string
	TaxValue int
}

// TaxLevelData Tax Level Struct for get tax level from database table master_tax_level
type TaxLevelData struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}
