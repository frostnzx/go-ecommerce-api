package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Name         string
	IsAdmin      bool
	CreatedAt    time.Time
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
