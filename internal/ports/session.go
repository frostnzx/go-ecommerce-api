package ports

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
	"github.com/google/uuid"
)

type SessionRepo interface {
	CreateSession(ctx context.Context, s *session.Session) (*session.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (*session.Session, error)
	RevokeSession(ctx context.Context, id uuid.UUID) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
}
