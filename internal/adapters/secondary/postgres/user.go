package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
	"github.com/google/uuid"
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
	_, err := ur.db.NamedExecContext(ctx, "INSERT INTO users (email, password_hash, name , is_admin) VALUES (:email, :password_hash, :name, :is_admin)", u)
	if err != nil {
		return fmt.Errorf("error create user %w", err)
	}
	return nil
}
func (ur *UserRepo) GetUser(ctx context.Context, email string) (*user.User, error) {
	var u user.User
	err := ur.db.GetContext(ctx, &u, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &u, nil
}
func (ur *UserRepo) GetUserById(ctx context.Context, id uuid.UUID) (*user.User, error) {
	var u user.User
	err := ur.db.GetContext(ctx, &u, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return &u, nil
}
func (ur *UserRepo) ListUsers(ctx context.Context) ([]*user.User, error) {
	var u []*user.User
	err := ur.db.SelectContext(ctx, &u, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	return u, nil
}
func (ur *UserRepo) UpdateUser(ctx context.Context, u user.User) error {
	_, err := ur.db.NamedExecContext(ctx, "UPDATE users SET name=:name , email=:email , password_hash=:password_hash , is_admin=:is_admin , created_at=:created_at", u)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}
func (ur *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := ur.db.ExecContext(ctx, "DELETE users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
