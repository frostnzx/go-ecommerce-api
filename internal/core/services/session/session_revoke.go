package session

import (
	"context"

	"github.com/google/uuid"
)

type RevokeSessionRequest struct {
	ID uuid.UUID
}

type RevokeSessionResponse struct {
	Success bool
}

func (s *Service) RevokeSession(ctx context.Context, id uuid.UUID) error {
	return s.sessionRepo.RevokeSession(ctx, id)
}
