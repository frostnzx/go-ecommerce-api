package user

import "context"

type GetUserProfileReq struct {
}
type GetUserProfileResp struct {
}

func (s *Service) GetUserProfile(ctx context.Context, req GetUserProfileReq) (*GetUserProfileResp, error) {
	
}
