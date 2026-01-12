package postgres

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/address"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AddressRepo struct {
	db *sqlx.DB
}

func NewAddressRepo(db *sqlx.DB) (*AddressRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &AddressRepo{db: db}, nil
}

func (ur *AddressRepo) Create(ctx context.Context, a address.Address) (address.Address, error) {
}
func (ur *AddressRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]address.Address, error) {
}
func (ur *AddressRepo) GetByID(ctx context.Context, id uuid.UUID) (address.Address, error) {
}
func (ur *AddressRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
}

func (ur *AddressRepo) SetDefault(ctx context.Context, userID, addressID uuid.UUID) error {
}
func (ur *AddressRepo) GetDefault(ctx context.Context, userID uuid.UUID) (address.Address, error) {
}
