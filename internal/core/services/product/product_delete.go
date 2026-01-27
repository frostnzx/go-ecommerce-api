package product

import (
	"context"
)

func (s *Service) DeleteProduct(ctx context.Context, req DeleteProductReq) error {
	// Check if product exists
	_, err := s.productRepo.GetByID(ctx, req.ID)
	if err != nil {
		return ErrProductNotFound
	}

	return s.productRepo.DeleteById(ctx, req.ID)
}
