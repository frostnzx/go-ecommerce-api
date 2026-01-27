package product

import (
	"context"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

func (s *Service) GetProduct(ctx context.Context, req GetProductReq) (*GetProductResp, error) {
	p, err := s.productRepo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrProductNotFound
	}

	return &GetProductResp{
		ID:          p.ID,
		SKU:         p.SKU,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		StockQty:    p.StockQty,
		Active:      p.Active,
		CreatedAt:   p.CreatedAt,
	}, nil
}

func (s *Service) ListProducts(ctx context.Context) (*ListProductsResp, error) {
	products, err := s.productRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	var productInfos []ProductInfo
	for _, p := range products {
		productInfos = append(productInfos, ProductInfo{
			ID:          p.ID,
			SKU:         p.SKU,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			StockQty:    p.StockQty,
			Active:      p.Active,
		})
	}

	return &ListProductsResp{
		Products: productInfos,
	}, nil
}
