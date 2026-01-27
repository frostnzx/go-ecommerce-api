package order

import (
	"context"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
	"github.com/google/uuid"
)

type API interface {
	PlaceOrder(context.Context, PlaceOrderReq) (*PlaceOrderResp, error)
	ListOrders(context.Context, ListOrdersReq) (*ListOrdersResp, error)
	GetOrder(context.Context, GetOrderReq) (*GetOrderResp, error)
	CancelOrder(context.Context, CancelOrderReq) error
	UpdateOrderStatus(context.Context, UpdateOrderStatusReq) error // Admin only
}

type Service struct {
	orderRepo   ports.OrderRepo
	itemsRepo   ports.ItemsRepo
	productRepo ports.ProductRepo
}

func NewService(or ports.OrderRepo, ir ports.ItemsRepo, pr ports.ProductRepo) *Service {
	return &Service{
		orderRepo:   or,
		itemsRepo:   ir,
		productRepo: pr,
	}
}

// Request/Response types

type OrderItemReq struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type PlaceOrderReq struct {
	UserID    uuid.UUID      `json:"user_id"`
	AddressID uuid.UUID      `json:"address_id"`
	Items     []OrderItemReq `json:"items"`
}

type PlaceOrderResp struct {
	ID          uuid.UUID `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListOrdersReq struct {
	UserID uuid.UUID `json:"user_id"`
}

type OrderInfo struct {
	ID          uuid.UUID `json:"id"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListOrdersResp struct {
	Orders []OrderInfo `json:"orders"`
}

type GetOrderReq struct {
	OrderID uuid.UUID `json:"order_id"`
	UserID  uuid.UUID `json:"user_id"` // For ownership verification
}

type OrderItemInfo struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

type GetOrderResp struct {
	ID          uuid.UUID       `json:"id"`
	UserID      uuid.UUID       `json:"user_id"`
	AddressID   uuid.UUID       `json:"address_id"`
	Status      string          `json:"status"`
	TotalAmount float64         `json:"total_amount"`
	Items       []OrderItemInfo `json:"items"`
	CreatedAt   time.Time       `json:"created_at"`
}

type CancelOrderReq struct {
	OrderID uuid.UUID `json:"order_id"`
	UserID  uuid.UUID `json:"user_id"` // For ownership verification
}

type UpdateOrderStatusReq struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"status"`
}
