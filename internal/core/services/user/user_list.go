package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type ListUsersResp struct {
	Users []UserInfo
}

type UserInfo struct {
	ID      uuid.UUID
	Name    string
	Email   string
	IsAdmin bool
}

func (s *Service) ListUsers(ctx context.Context) (*ListUsersResp, error) {
	users, err := s.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	var userInfos []UserInfo
	for _, u := range users {
		userInfos = append(userInfos, UserInfo{
			ID:      u.ID,
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: u.IsAdmin,
		})
	}

	return &ListUsersResp{Users: userInfos}, nil
}
