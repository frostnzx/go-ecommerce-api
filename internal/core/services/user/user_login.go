package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/session"
	"github.com/frostnzx/go-ecommerce-api/internal/core/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginUserResp struct {
	SessionID             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	IsAdmin               bool      `json:"is_admin"`
}

func (s *Service) LoginUser(ctx context.Context, req LoginUserReq) (*LoginUserResp, error) {
	user, err := s.userRepo.GetUser(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("Error getting user:%w", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil { // wrong password
		return nil, fmt.Errorf("Error password not matching:%w", err)
	}
	var secretKey = os.Getenv("JWT_SECRET")
	tokenMaker := utils.NewJWTMaker(secretKey)
	// create jwt
	accessToken, accessClaims, err := tokenMaker.CreateToken(user.ID, user.Email, user.IsAdmin, 15*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("Error creating accessToken:%w", err)
	}
	refreshToken, refreshClaims, err := tokenMaker.CreateToken(user.ID, user.Email, user.IsAdmin, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("Error creating refreshToken:%w", err)
	}

	// Create session using the session service
	sess := &session.Session{
		Email:        user.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		CreatedAt:    time.Now().UTC(),
		ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time,
	}
	createdSession, err := s.sessionService.CreateSession(ctx, sess)
	if err != nil {
		return nil, fmt.Errorf("Error creating session:%w", err)
	}

	res := LoginUserResp{
		SessionID:             createdSession.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		Name:                  user.Name,
		Email:                 user.Email,
		IsAdmin:               user.IsAdmin,
	}
	return &res, nil
}
