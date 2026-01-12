package order

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	PlaceOrder(context.Context, PlaceOrderReq) (*PlaceOrderResp, error)
	ListOrders(context.Context, ListOrdersReq) (*ListOrdersResp, error)
	GetOrder(context.Context, GetOrderReq) (*GetOrderResp, error)          // owner only
	CancelOrder(context.Context, CancelOrderReq) (*CancelOrderResp, error) // owner only
}

type Service struct {
	orderRepo ports.OrderRepo
}

func NewService(or ports.OrderRepo) *Service {
	return &Service{
		orderRepo: or,
	}
}
