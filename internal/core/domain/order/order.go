package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID
	UserId      uuid.UUID
	AddressID   uuid.UUID
	Status      OrderStatus
	TotalAmount float64
	CreatedAt   time.Time
}

func New(userId, addressId uuid.UUID, status OrderStatus, totalAmount float64) Order {
	return Order{
		ID:          uuid.New(),
		UserId:      userId,
		AddressID:   addressId,
		Status:      status,
		TotalAmount: totalAmount,
		CreatedAt:   time.Now().UTC(),
	}
}
