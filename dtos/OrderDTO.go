package dtos

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderDTO struct {
	ID    uuid.UUID
	Name  string
	Price decimal.Decimal
}
