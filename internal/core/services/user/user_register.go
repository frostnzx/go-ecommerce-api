package user

import (
	"context"
	"fmt"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
	"github.com/google/uuid"
)

type RegisterUserReq struct {
	Name         string
	Email        string
	PasswordHash string
	IsAdmin      bool
}
type RegisterUserResp struct {
	ID string
}

func (s *Service) RegisterUser(ctx context.Context, req RegisterUserReq) (*RegisterUserResp, error) {
	user := user.New(req.Email, req.PasswordHash, req.Name, req.IsAdmin)
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("fail to register user:%w", err)
	}
	return &RegisterUserResp{ID: user.ID.String()}, nil
}
