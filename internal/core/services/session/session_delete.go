package session

import (
	"context"

	"github.com/google/uuid"
)

type DeleteSessionRequest struct {
	ID uuid.UUID
}

type DeleteSessionResponse struct {
	Success bool
}

func (s *Service) DeleteSession(ctx context.Context, id uuid.UUID) error {
	return s.sessionRepo.DeleteSession(ctx, id)
}
