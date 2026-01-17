package session

import (
	"context"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
)

type CreateSessionRequest struct {
	Email        string
	RefreshToken string
	ExpiresAt    time.Time
}

type CreateSessionResponse struct {
	Session *session.Session
}

func (s *Service) CreateSession(ctx context.Context, sess *session.Session) (*session.Session, error) {
	return s.sessionRepo.CreateSession(ctx, sess)
}
