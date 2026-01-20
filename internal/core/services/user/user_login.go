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
	Email    string
	Password string
}
type LoginUserResp struct {
	SessionID             string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
	Name                  string
	Email                 string
	IsAdmin               bool
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

	// Generate sessionID
	sessionID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("error generating session ID:%w", err)
	}
	// Create refresh token FIRST
	refreshToken, refreshClaims, err := tokenMaker.CreateToken(sessionID.String(), user.ID, user.Email, user.IsAdmin, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("Error creating refreshToken:%w", err)
	}

	// Create access token with the SAME session ID
	accessToken, accessClaims, err := tokenMaker.CreateToken(sessionID.String(), user.ID, user.Email, user.IsAdmin, 15*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("Error creating accessToken:%w", err)
	}

	// Create session using the session service
	// Session ID matches both tokens' SessionID field
	sess := &session.Session{
		ID:           refreshClaims.SessionID,
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
