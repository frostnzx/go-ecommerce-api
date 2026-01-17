package user

import (
	"context"
	"fmt"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserReq struct {
	Name     string
	Email    string
	Password string
	IsAdmin  bool
}
type RegisterUserResp struct {
	ID uuid.UUID
}

func (s *Service) RegisterUser(ctx context.Context, req RegisterUserReq) (*RegisterUserResp, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("fail to hash password:%w", err)
	}
	user := user.New(req.Email, string(passwordHash), req.Name, req.IsAdmin)
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("fail to register user:%w", err)
	}
	resp := RegisterUserResp{
		ID: user.ID,
	}
	return &resp, nil
}
