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
	query := `
		INSERT INTO orders (id, user_id, address_id, status, total_amount, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, address_id, status, total_amount, created_at
	`
	var created order.Order
	err := or.db.QueryRowxContext(ctx, query,
		o.ID, o.UserId, o.AddressID, o.Status, o.TotalAmount, o.CreatedAt,
	).StructScan(&created)
	if err != nil {
		return order.Order{}, err
	}
	return created, nil
}

func (or *OrderRepo) GetByID(ctx context.Context, orderID uuid.UUID) (order.Order, error) {
	query := `
		SELECT id, user_id, address_id, status, total_amount, created_at
		FROM orders
		WHERE id = $1
	`
	var o order.Order
	err := or.db.GetContext(ctx, &o, query, orderID)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}

func (or *OrderRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]order.Order, error) {
	query := `
		SELECT id, user_id, address_id, status, total_amount, created_at
		FROM orders
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	var orders []order.Order
	err := or.db.SelectContext(ctx, &orders, query, userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepo) UpdateStatus(ctx context.Context, orderID uuid.UUID, status order.OrderStatus) error {
	query := `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`
	_, err := or.db.ExecContext(ctx, query, status, orderID)
	return err
}
