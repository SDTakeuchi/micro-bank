package model

type Currency struct {
	ID           string
	Code         string
	Name         string
	DecimalPoint uint8 // e.g.) USD: 2, JPY: 0
}
