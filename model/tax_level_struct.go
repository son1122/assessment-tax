package model

type TaxLevel struct {
	ID       int
	Floor    int
	Ceil     int
	CreateAt string
	TaxValue int
}
