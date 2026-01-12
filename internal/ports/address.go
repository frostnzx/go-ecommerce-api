package ports

import (
	"context"
	//"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/address"
	"github.com/google/uuid"
)

var (
	/* Errors */
)

type AddressRepo interface {
	Create(ctx context.Context, a address.Address) (address.Address, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]address.Address, error)
	GetByID(ctx context.Context, id uuid.UUID) (address.Address, error)
	DeleteById(ctx context.Context, id uuid.UUID) (error)

	SetDefault(ctx context.Context, userID, addressID uuid.UUID) error
	GetDefault(ctx context.Context, userID uuid.UUID) (address.Address, error)
}