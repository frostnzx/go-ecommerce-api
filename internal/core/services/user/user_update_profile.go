package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UpdateUserProfileReq struct {
	ID    uuid.UUID // Extracted from JWT claims in handler/middleware
	Name  string
	Email string
}

func (s *Service) UpdateUserProfile(ctx context.Context, req UpdateUserProfileReq) error {
	user, err := s.userRepo.GetUserByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	user.Name = req.Name
	user.Email = req.Email

	err = s.userRepo.UpdateUser(ctx, *user)
	if err != nil {
		return fmt.Errorf("error updating user profile: %w", err)
	}

	return nil
}
