package product

import (
	"context"
	"errors"

	"github.com/frostnzx/go-ecommerce-api/internal/core/domain/product"
)

var (
	ErrInvalidSKU   = errors.New("SKU is required")
	ErrInvalidName  = errors.New("name is required")
	ErrInvalidPrice = errors.New("price must be greater than 0")
)

func (s *Service) AddProduct(ctx context.Context, req AddProductReq) (*AddProductResp, error) {
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

	// Create product
	newProduct := product.New(req.SKU, req.Name, req.Description, req.Price, req.StockQty)

	created, err := s.productRepo.Create(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return &AddProductResp{
		ID:        created.ID,
		SKU:       created.SKU,
		Name:      created.Name,
		CreatedAt: created.CreatedAt,
	}, nil
}
