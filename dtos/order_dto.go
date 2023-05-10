package dtos

import (
	"github.com/shopspring/decimal"
)

type OrderDTO struct {
	ID    string
	Name  string
	Price decimal.Decimal
}
