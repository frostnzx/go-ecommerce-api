package address

import (
	"context"
	"errors"
)

func (s *Service) SetDefaultAddress(ctx context.Context, req SetDefaultAddressReq) error {
	// Verify ownership
	addr, err := s.addressRepo.GetByID(ctx, req.AddressID)
	if err != nil {
		return err
	}
	if addr.UserID != req.UserID {
		return errors.New("not authorized to modify this address")
	}

	return s.addressRepo.SetDefault(ctx, req.UserID, req.AddressID)
}
