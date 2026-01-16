package user

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	RegisterUser(context.Context, RegisterUserReq) (*RegisterUserResp, error)
	AuthenticateUser(context.Context, AuthenticateUserReq) (*AuthenticateUserResp, error)
	GetUserProfile(context.Context, GetUserProfileReq) (*GetUserProfileResp, error)
	UpdateUserProfile(context.Context, UpdateUserProfileReq) error
	ChangePassword(context.Context, ChangePasswordProfileReq) error
	DeleteAccount(context.Context, DeleteAccountProfileReq) error
}

type Service struct {
	userRepo ports.UserRepo
}

func NewService(ur ports.UserRepo) *Service {
	return &Service{
		userRepo: ur,
	}
}
