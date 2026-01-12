package address

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	AddAddress(context.Context, AddAddressReq) (*AddAddressResp, error)
	DeleteAddress(context.Context, DeleteAddressReq) (*DeleteAddressResp , error)
	ListAddresses(context.Context, ListAddressesReq) (*ListAddressesResp, error)
	SetDefaultAddress(context.Context, SetDefaultAddressReq) (*SetDefaultAddressResp, error)
	GetDefaultAddress(context.Context, GetDefaultAddressReq) (*GetDefaultAddressResp, error)
}

type Service struct {
	addressRepo ports.AddressRepo
}

func NewService(ar ports.AddressRepo) *Service {
	return &Service{
		addressRepo: ar,
	}
}
