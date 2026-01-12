package ports

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/order"
	"github.com/google/uuid"
)

var (
	ErrCreateOrder = errors.New("cannot create order")
	ErrAddItem = errors.New("cannot add item")
)

type OrderRepo interface {
	Create(ctx context.Context, o order.Order) (order.Order, error)
	AddItems(ctx context.Context, orderID uuid.UUID, items []order.Order) error

	GetByID(ctx context.Context, orderID uuid.UUID) (order.Order, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]order.Order, error)

	UpdateStatus(ctx context.Context, orderID uuid.UUID, status order.OrderStatus) error
}
