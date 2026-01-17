package session

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
	"github.com/google/uuid"
)

type GetSessionRequest struct {
	ID uuid.UUID
}

type GetSessionResponse struct {
	Session *session.Session
}

func (s *Service) GetSession(ctx context.Context, id uuid.UUID) (*session.Session, error) {
	return s.sessionRepo.GetSession(ctx, id)
}
