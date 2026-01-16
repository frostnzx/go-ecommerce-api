package items

import (
	"github.com/google/uuid"
)

type Items struct {
	ID                uuid.UUID `db:id`
	OrderID           uuid.UUID `db:order_id`
	ProductID         uuid.UUID `db:product_id`
	Quantity          int       `db:quantity`
	UnitPriceSnapshot float64   `db:unit_price_snapshot` // price at that exact moment why click buy
}

func New(orderId, productId uuid.UUID, quantity int, unitPrice float64) Items {
	return Items{
		ID:                uuid.New(),
		OrderID:           orderId,
		ProductID:         productId,
		Quantity:          quantity,
		UnitPriceSnapshot: unitPrice,
	}
}
