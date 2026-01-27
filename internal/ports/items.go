package ports

import (
	"context"
	//"errors"
	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/items"
	"github.com/google/uuid"
)

var (
/* Errors */
)

type ItemsRepo interface {
	Create(ctx context.Context, item items.Items) (items.Items, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]items.Items, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]items.Items, error)
	GetByID(ctx context.Context, id uuid.UUID) (items.Items, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
}
