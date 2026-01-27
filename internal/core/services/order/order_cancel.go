package order

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/order"
)

var (
	ErrCannotCancelOrder = errors.New("order cannot be cancelled in current status")
)

func (s *Service) CancelOrder(ctx context.Context, req CancelOrderReq) error {
	// Get the order
	o, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return ErrOrderNotFound
	}

	// Verify ownership
	if o.UserId != req.UserID {
		return ErrNotOrderOwner
	}

	// Only pending orders can be cancelled
	if o.Status != order.OrderPending {
		return ErrCannotCancelOrder
	}

	return s.orderRepo.UpdateStatus(ctx, req.OrderID, order.OrderCancelled)
}
