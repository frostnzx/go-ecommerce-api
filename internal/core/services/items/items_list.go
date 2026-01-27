// filepath: /home/frostnzx/dev/go-ecommerce-api/internal/core/services/items/items_list.go
package items

import (
	"context"
)

func (s *Service) GetItem(ctx context.Context, req GetItemReq) (*GetItemResp, error) {
	item, err := s.itemsRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrItemNotFound
	}

	return &GetItemResp{
		ID:                item.ID,
		OrderID:           item.OrderID,
		ProductID:         item.ProductID,
		Quantity:          item.Quantity,
		UnitPriceSnapshot: item.UnitPriceSnapshot,
	}, nil
}

func (s *Service) ListItemsByOrder(ctx context.Context, req ListItemsByOrderReq) (*ListItemsByOrderResp, error) {
	// Verify order ownership
	order, err := s.orderRepo.GetByID(ctx, req.OrderID)
	if err != nil {
		return nil, ErrOrderNotFound
	}
	if order.UserId != req.UserID {
		return nil, ErrNotOrderOwner
	}

	items, err := s.itemsRepo.ListByOrderID(ctx, req.OrderID)
	if err != nil {
		return nil, err
	}

	var itemInfos []ItemInfo
	for _, item := range items {
		itemInfos = append(itemInfos, ItemInfo{
			ID:                item.ID,
			ProductID:         item.ProductID,
			Quantity:          item.Quantity,
			UnitPriceSnapshot: item.UnitPriceSnapshot,
		})
	}

	return &ListItemsByOrderResp{
		Items: itemInfos,
	}, nil
}

func (s *Service) ListItemsByUser(ctx context.Context, req ListItemsByUserReq) (*ListItemsByUserResp, error) {
	items, err := s.itemsRepo.ListByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	var itemInfos []ItemInfo
	for _, item := range items {
		itemInfos = append(itemInfos, ItemInfo{
			ID:                item.ID,
			ProductID:         item.ProductID,
			Quantity:          item.Quantity,
			UnitPriceSnapshot: item.UnitPriceSnapshot,
		})
	}

	return &ListItemsByUserResp{
		Items: itemInfos,
	}, nil
}
