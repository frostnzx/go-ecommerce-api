package postgres

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/items"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ItemsRepo struct {
	db *sqlx.DB
}

func NewItemsRepo(db *sqlx.DB) (*ItemsRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &ItemsRepo{db: db}, nil
}

func (ir *ItemsRepo) Create(ctx context.Context, item items.Items) (items.Items, error) {
	query := `
		INSERT INTO items (id, order_id, product_id, quantity, unit_price_snapshot)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, order_id, product_id, quantity, unit_price_snapshot
	`
	var created items.Items
	err := ir.db.QueryRowxContext(ctx, query,
		item.ID, item.OrderID, item.ProductID, item.Quantity, item.UnitPriceSnapshot,
	).StructScan(&created)
	if err != nil {
		return items.Items{}, err
	}
	return created, nil
}

func (ir *ItemsRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]items.Items, error) {
	query := `
		SELECT i.id, i.order_id, i.product_id, i.quantity, i.unit_price_snapshot
		FROM items i
		JOIN orders o ON i.order_id = o.id
		WHERE o.user_id = $1
	`
	var itemsList []items.Items
	err := ir.db.SelectContext(ctx, &itemsList, query, userID)
	if err != nil {
		return nil, err
	}
	return itemsList, nil
}

func (ir *ItemsRepo) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]items.Items, error) {
	query := `
		SELECT id, order_id, product_id, quantity, unit_price_snapshot
		FROM items
		WHERE order_id = $1
	`
	var itemsList []items.Items
	err := ir.db.SelectContext(ctx, &itemsList, query, orderID)
	if err != nil {
		return nil, err
	}
	return itemsList, nil
}

func (ir *ItemsRepo) GetByID(ctx context.Context, id uuid.UUID) (items.Items, error) {
	query := `
		SELECT id, order_id, product_id, quantity, unit_price_snapshot
		FROM items
		WHERE id = $1
	`
	var item items.Items
	err := ir.db.GetContext(ctx, &item, query, id)
	if err != nil {
		return items.Items{}, err
	}
	return item, nil
}

func (ir *ItemsRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := ir.db.ExecContext(ctx, query, id)
	return err
}
