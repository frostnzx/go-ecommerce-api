package user

import (
	"time"

	"github.com/google/uuid"
)

type registerUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}
type registerUserResp struct {
	ID uuid.UUID `json:"id"`
}

type loginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginUserResp struct {
	SessionID             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	IsAdmin               bool      `json:"is_admin"`
}

type logoutUserReq struct {
	SessionID string `json:"session_id"`
}

type renewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
type renewAccessTokenResp struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type updateUserProfileReq struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type getUserProfileReq struct {
	ID uuid.UUID `json:"id"`
}

type getUserProfileResp struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	IsAdmin bool      `json:"is_admin"`
}

type deleteAccountProfileReq struct {
	ID        uuid.UUID `json:"id"`
	SessionID string    `json:"session_id"`
}

type changePasswordProfileReq struct {
	ID          uuid.UUID `json:"id"`
	CurrentPassword string 
	NewPassword     string
}