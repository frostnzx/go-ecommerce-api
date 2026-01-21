package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type GetUserProfileReq struct {
	UserID uuid.UUID // Extracted from JWT claims in handler/middleware
}

type GetUserProfileResp struct {
	ID      uuid.UUID
	Name    string
	Email   string
	IsAdmin bool
}

func (s *Service) GetUserProfile(ctx context.Context, req GetUserProfileReq) (*GetUserProfileResp, error) {
	user, err := s.userRepo.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &GetUserProfileResp{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
	}, nil
}
