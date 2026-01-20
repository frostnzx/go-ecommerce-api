package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/core/utils"
)

type RenewAccessTokenReq struct {
	RefreshToken string
}

type RenewAccessTokenResp struct {
	AccessToken          string
	AccessTokenExpiresAt time.Time
}

func (s *Service) RenewAccessToken(ctx context.Context, req RenewAccessTokenReq) (*RenewAccessTokenResp, error) {
	var secretKey = os.Getenv("JWT_SECRET")
	tokenMaker := utils.NewJWTMaker(secretKey)

	refreshClaims, err := tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying token")
	}

	// Use SessionID (shared between access and refresh tokens)
	session, err := s.sessionService.GetSession(ctx, refreshClaims.SessionID)
	if err != nil {
		return nil, fmt.Errorf("error getting session:%w", err)
	}
	if session.IsRevoked {
		return nil, fmt.Errorf("session revoked")
	}
	if session.Email != refreshClaims.Email {
		return nil, fmt.Errorf("invalid session")
	}

	// Create new access token with the same session ID
	accessToken, accessClaims, err := tokenMaker.CreateToken(
		refreshClaims.SessionID,
		refreshClaims.ID,
		refreshClaims.Email,
		refreshClaims.IsAdmin,
		15*time.Minute,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating access token:%w", err)
	}

	return &RenewAccessTokenResp{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	}, nil
}

/*
	- VerifyToken gets refreshClaims from the refreshToken string
	- GetSession use sessionID inside refreshClaims to get a session object
	- Create new accessToken with properties from refreshClaims
*/
