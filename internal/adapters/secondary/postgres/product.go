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

func (pr *ProductRepo) Create(ctx context.Context, item product.Product) (product.Product, error) {
	
}
func (pr *ProductRepo) GetByID(ctx context.Context, id uuid.UUID) (product.Product, error) {

}
func (pr *ProductRepo) DeleteById(ctx context.Context, id uuid.UUID) error {

}
func (pr *ProductRepo) EditById(ctx context.Context, id uuid.UUID) (product.Product, error) {

}
