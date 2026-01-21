package session

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	CreateSession(ctx context.Context, s *session.Session) (*session.Session, error)
	GetSession(ctx context.Context, id string) (*session.Session, error)
	RevokeSession(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error
}

type Service struct {
	sessionRepo ports.SessionRepo
}

func NewService(sr ports.SessionRepo) *Service {
	return &Service{
		sessionRepo: sr,
	}
}



