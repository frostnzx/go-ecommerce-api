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

func (ar *AddressRepo) Create(ctx context.Context, a address.Address) (address.Address, error) {
	query := `
		INSERT INTO addresses (id, user_id, line1, city, province, postal_code, country, is_default)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, user_id, line1, city, province, postal_code, country, is_default
	`
	var created address.Address
	err := ar.db.QueryRowxContext(ctx, query,
		a.ID, a.UserID, a.Line1, a.City, a.Province, a.PostalCode, a.Country, a.IsDefault,
	).StructScan(&created)
	if err != nil {
		return address.Address{}, err
	}
	return created, nil
}

func (ar *AddressRepo) ListByUserID(ctx context.Context, userID uuid.UUID) ([]address.Address, error) {
	query := `
		SELECT id, user_id, line1, city, province, postal_code, country, is_default
		FROM addresses
		WHERE user_id = $1
		ORDER BY is_default DESC, id
	`
	var addresses []address.Address
	err := ar.db.SelectContext(ctx, &addresses, query, userID)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (ar *AddressRepo) GetByID(ctx context.Context, id uuid.UUID) (address.Address, error) {
	query := `
		SELECT id, user_id, line1, city, province, postal_code, country, is_default
		FROM addresses
		WHERE id = $1
	`
	var a address.Address
	err := ar.db.GetContext(ctx, &a, query, id)
	if err != nil {
		return address.Address{}, err
	}
	return a, nil
}

func (ar *AddressRepo) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM addresses WHERE id = $1`
	_, err := ar.db.ExecContext(ctx, query, id)
	return err
}

func (ar *AddressRepo) SetDefault(ctx context.Context, userID, addressID uuid.UUID) error {
	// First, unset all defaults for this user
	unsetQuery := `UPDATE addresses SET is_default = false WHERE user_id = $1`
	_, err := ar.db.ExecContext(ctx, unsetQuery, userID)
	if err != nil {
		return err
	}

	// Then set the new default
	setQuery := `UPDATE addresses SET is_default = true WHERE id = $1 AND user_id = $2`
	_, err = ar.db.ExecContext(ctx, setQuery, addressID, userID)
	return err
}

func (ar *AddressRepo) GetDefault(ctx context.Context, userID uuid.UUID) (address.Address, error) {
	query := `
		SELECT id, user_id, line1, city, province, postal_code, country, is_default
		FROM addresses
		WHERE user_id = $1 AND is_default = true
	`
	var a address.Address
	err := ar.db.GetContext(ctx, &a, query, userID)
	if err != nil {
		return address.Address{}, err
	}
	return a, nil
}
