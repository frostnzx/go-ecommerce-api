package session

import (
	"time"
)

type Session struct {
	ID           string    `db:"id"`
	Email        string    `db:"user_email"`
	RefreshToken string    `db:"refresh_token"`
	IsRevoked    bool      `db:"is_revoked"`
	CreatedAt    time.Time `db:"created_at"`
	ExpiresAt    time.Time `db:"expires_at"`
}

func New(id string, email, refreshToken string, isRevoked bool, exp time.Time) Session {
	return Session{
		ID:           id,
		Email:        email,
		RefreshToken: refreshToken,
		IsRevoked:    isRevoked,
		CreatedAt:    time.Now().UTC(),
		ExpiresAt:    exp,
	}
}
