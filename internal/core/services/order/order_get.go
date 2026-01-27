package order

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/order"
)

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrNotOrderOwner = errors.New("not authorized to access this order")
)

func (s *Service) GetOrder(ctx context.Context, req GetOrderReq) (*GetOrderResp, error) {
	// Get the order
	o, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, ErrOrderNotFound
	}

	// Verify ownership
	if o.UserId != req.UserID {
		return nil, ErrNotOrderOwner
	}

	// Get order items
	orderItems, err := s.itemsRepo.ListByOrderID(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	var itemInfos []OrderItemInfo
	for _, item := range orderItems {
		itemInfos = append(itemInfos, OrderItemInfo{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPriceSnapshot,
		})
	}

	return &GetOrderResp{
		ID:          o.ID,
		UserID:      o.UserId,
		AddressID:   o.AddressID,
		Status:      string(o.Status),
		TotalAmount: o.TotalAmount,
		Items:       itemInfos,
		CreatedAt:   o.CreatedAt,
	}, nil
}

func (s *Service) ListOrders(ctx context.Context, req ListOrdersReq) (*ListOrdersResp, error) {
	orders, err := s.orderRepo.ListByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	var orderInfos []OrderInfo
	for _, o := range orders {
		orderInfos = append(orderInfos, OrderInfo{
			ID:          o.ID,
			Status:      string(o.Status),
			TotalAmount: o.TotalAmount,
			CreatedAt:   o.CreatedAt,
		})
	}

	return &ListOrdersResp{
		Orders: orderInfos,
	}, nil
}

func (s *Service) UpdateOrderStatus(ctx context.Context, req UpdateOrderStatusReq) error {
	// Verify the order exists
	_, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return ErrOrderNotFound
	}

	return s.orderRepo.UpdateStatus(ctx, req.OrderID, order.OrderStatus(req.Status))
}
