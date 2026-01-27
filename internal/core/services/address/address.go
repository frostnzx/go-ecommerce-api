package address

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
	"github.com/google/uuid"
)

type API interface {
	AddAddress(context.Context, AddAddressReq) (*AddAddressResp, error)
	DeleteAddress(context.Context, DeleteAddressReq) error
	ListAddresses(context.Context, ListAddressesReq) (*ListAddressesResp, error)
	SetDefaultAddress(context.Context, SetDefaultAddressReq) error
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

// Request/Response types

type AddAddressReq struct {
	UserID     uuid.UUID `json:"user_id"`
	Line1      string    `json:"line1"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
}

type AddAddressResp struct {
	ID uuid.UUID `json:"id"`
}

type DeleteAddressReq struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"` // For ownership verification
}

type ListAddressesReq struct {
	UserID uuid.UUID `json:"user_id"`
}

type AddressInfo struct {
	ID         uuid.UUID `json:"id"`
	Line1      string    `json:"line1"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	IsDefault  bool      `json:"is_default"`
}

type ListAddressesResp struct {
	Addresses []AddressInfo `json:"addresses"`
}

type SetDefaultAddressReq struct {
	AddressID uuid.UUID `json:"address_id"`
	UserID    uuid.UUID `json:"user_id"`
}

type GetDefaultAddressReq struct {
	UserID uuid.UUID `json:"user_id"`
}

type GetDefaultAddressResp struct {
	ID         uuid.UUID `json:"id"`
	Line1      string    `json:"line1"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
}
