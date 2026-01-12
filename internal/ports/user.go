package ports

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
)

var (
	ErrUserNotFound = errors.New("user does not exist")
	ErrCreateUser = errors.New("cannot create user")
)

type UserRepo interface {
	Create(ctx context.Context, u user.User) error
	GetByID(ctx context.Context, id int64) (user.User, error)
	GetByEmail(ctx context.Context, email string) (user.User, error)
}
