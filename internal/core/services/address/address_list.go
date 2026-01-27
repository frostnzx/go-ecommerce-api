package address

import (
	"context"
)

func (s *Service) ListAddresses(ctx context.Context, req ListAddressesReq) (*ListAddressesResp, error) {
	addresses, err := s.addressRepo.ListByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	var addressInfos []AddressInfo
	for _, a := range addresses {
		addressInfos = append(addressInfos, AddressInfo{
			ID:         a.ID,
			Line1:      a.Line1,
			City:       a.City,
			Province:   a.Province,
			PostalCode: a.PostalCode,
			Country:    a.Country,
			IsDefault:  a.IsDefault,
		})
	}

	return &ListAddressesResp{
		Addresses: addressInfos,
	}, nil
}
