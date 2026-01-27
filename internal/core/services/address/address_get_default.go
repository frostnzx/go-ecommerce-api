package address

import (
	"context"
)

func (s *Service) GetDefaultAddress(ctx context.Context, req GetDefaultAddressReq) (*GetDefaultAddressResp, error) {
	addr, err := s.addressRepo.GetDefault(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &GetDefaultAddressResp{
		ID:         addr.ID,
		Line1:      addr.Line1,
		City:       addr.City,
		Province:   addr.Province,
		PostalCode: addr.PostalCode,
		Country:    addr.Country,
	}, nil
}
