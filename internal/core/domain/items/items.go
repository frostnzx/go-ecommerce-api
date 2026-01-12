package items

import (
	"github.com/google/uuid"
)

type Items struct {
	ID                uuid.UUID
	OrderID           uuid.UUID
	ProductID         uuid.UUID
	Quantity          int
	UnitPriceSnapshot float64 // price at that exact moment why click buy
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
