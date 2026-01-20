package session

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
)

func (s *Service) GetSession(ctx context.Context, id string) (*session.Session, error) {
	return s.sessionRepo.GetSession(ctx, id)
}
