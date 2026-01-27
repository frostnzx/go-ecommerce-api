// filepath: /home/frostnzx/dev/go-ecommerce-api/internal/core/services/items/items_delete.go
package items

import (
	"context"
)

func (s *Service) DeleteItem(ctx context.Context, req DeleteItemReq) error {
	// Verify order ownership
	order, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return ErrOrderNotFound
	}
	if order.UserId != req.UserID {
		return ErrNotOrderOwner
	}

	// Verify item exists and belongs to the order
	item, err := s.itemsRepo.GetByID(ctx, req.ID)
	if err != nil {
		return ErrItemNotFound
	}
	if item.OrderID != req.OrderID {
		return ErrItemNotFound
	}

	return s.itemsRepo.DeleteById(ctx, req.ID)
}
