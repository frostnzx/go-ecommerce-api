package product

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	AddProduct(context.Context, AddProductReq) (*AddProductResp, error)
	GetProduct(context.Context, GetProductReq) (*GetProductResp, error)
	EditProduct(context.Context, EditProductReq) (*EditProductResp, error)
	DeleteProduct(context.Context, DeleteProductReq) (*DeleteProductResp, error)
}

type Service struct {
	productRepo ports.ProductRepo
}

func NewService(pr ports.ProductRepo) *Service {
	return &Service{
		productRepo: pr,
	}
}
