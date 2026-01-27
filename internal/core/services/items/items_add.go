package items

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/items"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than 0")
	ErrProductNotFound = errors.New("product not found")
	ErrOrderNotFound   = errors.New("order not found")
	ErrNotOrderOwner   = errors.New("not authorized to modify this order")
	ErrItemNotFound    = errors.New("item not found")
)

func (s *Service) AddItem(ctx context.Context, req AddItemReq) (*AddItemResp, error) {
	if req.Quantity <= 0 {
		return nil, ErrInvalidQuantity
	}

	// Verify order exists
	_, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, ErrOrderNotFound
	}

	// Get product to get current price
	product, err := s.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	// Create item with price snapshot
	newItem := items.New(req.OrderID, req.ProductID, req.Quantity, product.Price)

	created, err := s.itemsRepo.Create(ctx, newItem)
	if err != nil {
		return nil, err
	}

	return &AddItemResp{
		ID:                created.ID,
		OrderID:           created.OrderID,
		ProductID:         created.ProductID,
		Quantity:          created.Quantity,
		UnitPriceSnapshot: created.UnitPriceSnapshot,
	}, nil
}
