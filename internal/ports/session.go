package ports

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
)

type SessionRepo interface {
	CreateSession(ctx context.Context, s *session.Session) (*session.Session, error)
	GetSession(ctx context.Context, id string) (*session.Session, error)
	RevokeSession(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error
}
