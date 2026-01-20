package session

import (
	"context"
)

func (s *Service) RevokeSession(ctx context.Context, id string) error {
	return s.sessionRepo.RevokeSession(ctx, id)
}
