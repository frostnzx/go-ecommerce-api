package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	IsAdmin      bool      `db:"is_admin"`
	CreatedAt    time.Time `db:"created_at"`
}
type Tokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

func New(email, passwordHash, name string, isAdmin bool) User {
	return User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		IsAdmin:      isAdmin,
		CreatedAt:    time.Now().UTC(),
	}
}
