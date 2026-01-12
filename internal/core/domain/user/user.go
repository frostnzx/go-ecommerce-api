package user

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Name         string
	CreatedAt    time.Time
}

func New(email, passwordHash, name string) User {
	return User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		CreatedAt:    time.Now().UTC(),
	}
}
