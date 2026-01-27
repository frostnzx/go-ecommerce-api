package order

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/items"
	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/order"
	"github.com/google/uuid"
)

var (
	ErrEmptyOrder        = errors.New("order must have at least one item")
	ErrInvalidQuantity   = errors.New("quantity must be greater than 0")
	ErrProductNotFound   = errors.New("product not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)

func (s *Service) PlaceOrder(ctx context.Context, req PlaceOrderReq) (*PlaceOrderResp, error) {
	if len(req.Items) == 0 {
		return nil, ErrEmptyOrder
	}

	// Calculate total amount and validate products
	var totalAmount float64
	type itemWithPrice struct {
		ProductID uuid.UUID
		Quantity  int
		UnitPrice float64
	}
	var itemsWithPrices []itemWithPrice

	for _, item := range req.Items {
		if item.Quantity <= 0 {
			return nil, ErrInvalidQuantity
		}

		// Get product to validate and get price
		product, err := s.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, ErrProductNotFound
		}

		// Check stock
		if product.StockQty < item.Quantity {
			return nil, ErrInsufficientStock
		}

		itemsWithPrices = append(itemsWithPrices, itemWithPrice{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: product.Price,
		})

		totalAmount += product.Price * float64(item.Quantity)
	}

	// Create the order
	newOrder := order.New(req.UserID, req.AddressID, order.OrderPending, totalAmount)
	createdOrder, err := s.orderRepo.Create(ctx, newOrder)
	if err != nil {
		return nil, err
	}

	// Create order items
	for _, item := range itemsWithPrices {
		newItem := items.New(createdOrder.ID, item.ProductID, item.Quantity, item.UnitPrice)
		_, err := s.itemsRepo.Create(ctx, newItem)
		if err != nil {
			return nil, err
		}
	}

	return &PlaceOrderResp{
		ID:          createdOrder.ID,
		TotalAmount: createdOrder.TotalAmount,
		Status:      string(createdOrder.Status),
		CreatedAt:   createdOrder.CreatedAt,
	}, nil
}
