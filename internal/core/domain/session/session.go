package session

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	RefreshToken string    `db:"refresh_token"`
	IsRevoked    bool      `db:"is_revoked"`
	CreatedAt    time.Time `db:"created_at"`
	ExpiresAt    time.Time `db:"expires_at"`
}

func New(id uuid.UUID, email, refreshToken string, isRevoked bool, exp time.Time) Session {
	return Session{
		ID:           uuid.New(),
		Email:        email,
		RefreshToken: refreshToken,
		IsRevoked:    isRevoked,
		CreatedAt:    time.Now().UTC(),
		ExpiresAt:    exp,
	}
}
