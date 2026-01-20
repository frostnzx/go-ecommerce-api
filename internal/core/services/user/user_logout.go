package user

import (
	"context"
	"fmt"
)

type LogoutUserReq struct {
	SessionID string
}

func (s *Service) LogoutUser(ctx context.Context, req LogoutUserReq) error {
	id := req.SessionID // SessionID passed down in req by middleware layer
	err := s.sessionService.DeleteSession(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting session:%w", err)
	}
	return nil
}
