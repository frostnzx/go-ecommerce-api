package product

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/product"
)

func (s *Service) EditProduct(ctx context.Context, req EditProductReq) (*EditProductResp, error) {
	// Validate input
	if req.SKU == "" {
		return nil, ErrInvalidSKU
	}
	if req.Name == "" {
		return nil, ErrInvalidName
	}
	if req.Price <= 0 {
		return nil, ErrInvalidPrice
	}

	// Check if product exists
	existing, err := s.productRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	// Update product
	updated := product.Product{
		ID:          existing.ID,
		SKU:         req.SKU,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		StockQty:    req.StockQty,
		Active:      req.Active,
		CreatedAt:   existing.CreatedAt,
	}

	result, err := s.productRepo.UpdateById(ctx, updated)
	if err != nil {
		return nil, err
	}

	return &EditProductResp{
		ID:          result.ID,
		SKU:         result.SKU,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		StockQty:    result.StockQty,
		Active:      result.Active,
	}, nil
}
