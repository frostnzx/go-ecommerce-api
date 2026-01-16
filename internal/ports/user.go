package ports

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
	ErrCreateUser   = errors.New("cannot create user")
)

type UserRepo interface {
	Create(ctx context.Context, u user.User) error
	GetUser(ctx context.Context, email string) (*user.User, error)
	ListUsers(ctx context.Context) ([]*user.User, error)
	UpdateUser(ctx context.Context, u user.User) error
	DeleteUser(ctx context.Context, id int64) error
}
