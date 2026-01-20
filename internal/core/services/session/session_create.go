package session

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
)

func (s *Service) CreateSession(ctx context.Context, sess *session.Session) (*session.Session, error) {
	return s.sessionRepo.CreateSession(ctx, sess)
}
