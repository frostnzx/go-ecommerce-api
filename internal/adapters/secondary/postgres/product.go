package postgres

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/product"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) (*ProductRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &ProductRepo{db: db}, nil
}

func (pr *ProductRepo) Create(ctx context.Context, p product.Product) (product.Product, error) {
	query := `
		INSERT INTO products (id, sku, name, description, price, stock_qty, active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, sku, name, description, price, stock_qty, active, created_at
	`
	var created product.Product
	err := pr.db.QueryRowxContext(ctx, query,
		p.ID, p.SKU, p.Name, p.Description, p.Price, p.StockQty, p.Active, p.CreatedAt,
	).StructScan(&created)
	if err != nil {
		return product.Product{}, err
	}
	return created, nil
}

func (pr *ProductRepo) GetByID(ctx context.Context, id uuid.UUID) (product.Product, error) {
	query := `
		SELECT id, sku, name, description, price, stock_qty, active, created_at
		FROM products
		WHERE id = $1
	`
	var p product.Product
	err := pr.db.GetContext(ctx, &p, query, id)
	if err != nil {
		return product.Product{}, err
	}
	return p, nil
}

func (pr *ProductRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := pr.db.ExecContext(ctx, query, id)
	return err
}

func (pr *ProductRepo) UpdateById(ctx context.Context, p product.Product) (product.Product, error) {
	query := `
		UPDATE products
		SET sku = $1, name = $2, description = $3, price = $4, stock_qty = $5, active = $6
		WHERE id = $7
		RETURNING id, sku, name, description, price, stock_qty, active, created_at
	`
	var updated product.Product
	err := pr.db.QueryRowxContext(ctx, query,
		p.SKU, p.Name, p.Description, p.Price, p.StockQty, p.Active, p.ID,
	).StructScan(&updated)
	if err != nil {
		return product.Product{}, err
	}
	return updated, nil
}

func (pr *ProductRepo) List(ctx context.Context) ([]product.Product, error) {
	query := `
		SELECT id, sku, name, description, price, stock_qty, active, created_at
		FROM products
		WHERE active = true
		ORDER BY created_at DESC
	`
	var products []product.Product
	err := pr.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pr *ProductRepo) UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error {
	query := `UPDATE products SET stock_qty = $1 WHERE id = $2`
	_, err := pr.db.ExecContext(ctx, query, quantity, id)
	return err
}
