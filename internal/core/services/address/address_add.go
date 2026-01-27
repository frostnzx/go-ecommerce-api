package address

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/address"
)

var (
	ErrEmptyLine1   = errors.New("address line1 is required")
	ErrEmptyCity    = errors.New("city is required")
	ErrEmptyCountry = errors.New("country is required")
)

func (s *Service) AddAddress(ctx context.Context, req AddAddressReq) (*AddAddressResp, error) {
	if req.Line1 == "" {
		return nil, ErrEmptyLine1
	}
	if req.City == "" {
		return nil, ErrEmptyCity
	}
	if req.Country == "" {
		return nil, ErrEmptyCountry
	}

	newAddress := address.New(req.UserID, req.Line1, req.City, req.Province, req.PostalCode, req.Country)

	created, err := s.addressRepo.Create(ctx, newAddress)
	if err != nil {
		return nil, err
	}

	return &AddAddressResp{
		ID: created.ID,
	}, nil
}

func (s *Service) DeleteAddress(ctx context.Context, req DeleteAddressReq) error {
	// Verify ownership by getting the address first
	addr, err := s.addressRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if addr.UserID != req.UserID {
		return errors.New("not authorized to delete this address")
	}

	return s.addressRepo.DeleteById(ctx, req.ID)
}
