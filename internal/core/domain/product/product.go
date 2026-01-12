package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	SKU         string // stock keeping unit
	Name        string
	Description string
	Price       float64
	StockQty    int // stock quantity
	Active      bool
	CreatedAt   time.Time
}

func New(sku, name, desc string, price float64, stockQty int) Product {
	return Product{
		ID:          uuid.New(),
		SKU:         sku,
		Name:        name,
		Description: desc,
		Price:       price,
		StockQty:    stockQty,
		Active:      true,
		CreatedAt:   time.Now().UTC(),
	}
}
