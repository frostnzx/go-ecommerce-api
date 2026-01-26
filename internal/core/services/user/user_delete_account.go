package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type DeleteAccountProfileReq struct {
	ID        uuid.UUID // Extracted from JWT claims in handler/middleware
	SessionID string    // To delete the current session after account deletion
}

func (s *Service) DeleteAccount(ctx context.Context, req DeleteAccountProfileReq) error {
	err := s.userRepo.DeleteUser(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	err = s.sessionService.DeleteSession(ctx, req.SessionID)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}
