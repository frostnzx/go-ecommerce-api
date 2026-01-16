package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID   `db:id`
	UserId      uuid.UUID   `db:user_id`
	AddressID   uuid.UUID   `db:address_id`
	Status      OrderStatus `db:status`
	TotalAmount float64     `db:total_amount`
	CreatedAt   time.Time   `db:created_at`
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
