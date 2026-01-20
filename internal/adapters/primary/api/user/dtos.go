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
	ID uuid.UUID `json:id`
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
