package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordProfileReq struct {
	ID              uuid.UUID // Extracted from JWT claims in handler/middleware
	CurrentPassword string
	NewPassword     string
}

func (s *Service) ChangePassword(ctx context.Context, req ChangePasswordProfileReq) error {
	user, err := s.userRepo.GetUserByID(ctx, req.ID)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.CurrentPassword)); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing new password: %w", err)
	}

	user.PasswordHash = string(newHashedPassword)
	err = s.userRepo.UpdateUser(ctx, *user)
	if err != nil {
		return fmt.Errorf("error updating password: %w", err)
	}

	return nil
}
