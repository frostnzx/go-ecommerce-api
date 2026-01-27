package items

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
	"github.com/google/uuid"
)

type API interface {
	AddItem(ctx context.Context, req AddItemReq) (*AddItemResp, error)
	GetItem(ctx context.Context, req GetItemReq) (*GetItemResp, error)
	DeleteItem(ctx context.Context, req DeleteItemReq) error
	ListItemsByOrder(ctx context.Context, req ListItemsByOrderReq) (*ListItemsByOrderResp, error)
	ListItemsByUser(ctx context.Context, req ListItemsByUserReq) (*ListItemsByUserResp, error)
}

type Service struct {
	itemsRepo   ports.ItemsRepo
	productRepo ports.ProductRepo
	orderRepo   ports.OrderRepo
}

func NewService(ir ports.ItemsRepo, pr ports.ProductRepo, or ports.OrderRepo) *Service {
	return &Service{
		itemsRepo:   ir,
		productRepo: pr,
		orderRepo:   or,
	}
}

// Request/Response types

type AddItemReq struct {
	OrderID   uuid.UUID `json:"order_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

type AddItemResp struct {
	ID                uuid.UUID `json:"id"`
	OrderID           uuid.UUID `json:"order_id"`
	ProductID         uuid.UUID `json:"product_id"`
	Quantity          int       `json:"quantity"`
	UnitPriceSnapshot float64   `json:"unit_price_snapshot"`
}

type GetItemReq struct {
	ID uuid.UUID `json:"id"`
}

type GetItemResp struct {
	ID                uuid.UUID `json:"id"`
	OrderID           uuid.UUID `json:"order_id"`
	ProductID         uuid.UUID `json:"product_id"`
	Quantity          int       `json:"quantity"`
	UnitPriceSnapshot float64   `json:"unit_price_snapshot"`
}

type DeleteItemReq struct {
	ID      uuid.UUID `json:"id"`
	OrderID uuid.UUID `json:"order_id"` // For validation
	UserID  uuid.UUID `json:"user_id"`  // For ownership verification
}

type ListItemsByOrderReq struct {
	OrderID uuid.UUID `json:"order_id"`
	UserID  uuid.UUID `json:"user_id"` // For ownership verification
}

type ItemInfo struct {
	ID                uuid.UUID `json:"id"`
	ProductID         uuid.UUID `json:"product_id"`
	Quantity          int       `json:"quantity"`
	UnitPriceSnapshot float64   `json:"unit_price_snapshot"`
}

type ListItemsByOrderResp struct {
	Items []ItemInfo `json:"items"`
}

type ListItemsByUserReq struct {
	UserID uuid.UUID `json:"user_id"`
}

type ListItemsByUserResp struct {
	Items []ItemInfo `json:"items"`
}
