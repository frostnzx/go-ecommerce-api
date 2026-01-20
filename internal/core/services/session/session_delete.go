package session

import (
	"context"
)

func (s *Service) DeleteSession(ctx context.Context, id string) error {
	return s.sessionRepo.DeleteSession(ctx, id)
}
