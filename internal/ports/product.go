package ports

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/product"
	"github.com/google/uuid"
)

var (
/* Error */
)

type ProductRepo interface {
	Create(ctx context.Context, item product.Product) (product.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (product.Product, error)
	List(ctx context.Context) ([]product.Product, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	UpdateById(ctx context.Context, p product.Product) (product.Product, error)
	UpdateStock(ctx context.Context, id uuid.UUID, quantity int) error
}
