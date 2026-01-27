package user

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/services/session"
	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	RegisterUser(context.Context, RegisterUserReq) (*RegisterUserResp, error)
	GetUserProfile(context.Context, GetUserProfileReq) (*GetUserProfileResp, error)
	UpdateUserProfile(context.Context, UpdateUserProfileReq) error
	ChangePassword(context.Context, ChangePasswordProfileReq) error
	DeleteAccount(context.Context, DeleteAccountReq) error
	LoginUser(context.Context, LoginUserReq) (*LoginUserResp, error)
	LogoutUser(context.Context, LogoutUserReq) error
	RenewAccessToken(context.Context, RenewAccessTokenReq) (*RenewAccessTokenResp, error)
	ListUsers(context.Context) (*ListUsersResp, error) // Admin only
}

type Service struct {
	userRepo       ports.UserRepo
	sessionService session.API
}

func NewService(ur ports.UserRepo, ss session.API) *Service {
	return &Service{
		userRepo:       ur,
		sessionService: ss,
	}
}
