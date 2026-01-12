package postgres

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) (*UserRepo, error) {
	if db == nil {
		return nil, errors.New("database connection required")
	}
	return &UserRepo{db: db}, nil
}

func (ur *UserRepo) Create(ctx context.Context, u user.User) error {
}
func (ur *UserRepo) GetByID(ctx context.Context, id int64) (user.User, error) {
}
func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (user.User, error) {
}
