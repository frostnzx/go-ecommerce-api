package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `db:id`
	SKU         string    `db:sku` // stock keeping unit
	Name        string    `db:name`
	Description string    `db:description`
	Price       float64   `db:price`
	StockQty    int       `db:stock_qty` // stock quantity
	Active      bool      `db:active`
	CreatedAt   time.Time `db:created_at`
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
