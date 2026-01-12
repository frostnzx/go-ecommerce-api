package postgres

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/order"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) (*OrderRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &OrderRepo{db: db}, nil
}

func (or *OrderRepo) Create(ctx context.Context, o order.Order) (order.Order, error) {
}
func (or *OrderRepo) AddItems(ctx context.Context, orderID uuid.UUID, items []order.Order) error {
}

func (or *OrderRepo) GetByID(ctx context.Context, orderID uuid.UUID) (order.Order, error) {
}
func (or *OrderRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]order.Order, error) {
}

func (or *OrderRepo) UpdateStatus(ctx context.Context, orderID uuid.UUID, status order.OrderStatus) error {
}
