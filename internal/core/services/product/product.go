package product

import (
	"context"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
	"github.com/google/uuid"
)

type API interface {
	AddProduct(context.Context, AddProductReq) (*AddProductResp, error)
	GetProduct(context.Context, GetProductReq) (*GetProductResp, error)
	ListProducts(context.Context) (*ListProductsResp, error)
	EditProduct(context.Context, EditProductReq) (*EditProductResp, error)
	DeleteProduct(context.Context, DeleteProductReq) error
}

type Service struct {
	productRepo ports.ProductRepo
}

func NewService(pr ports.ProductRepo) *Service {
	return &Service{
		productRepo: pr,
	}
}

// Request/Response types

type AddProductReq struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
}

type AddProductResp struct {
	ID        uuid.UUID `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GetProductReq struct {
	ID uuid.UUID `json:"id"`
}

type GetProductResp struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductInfo struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
}

type ListProductsResp struct {
	Products []ProductInfo `json:"products"`
}

type EditProductReq struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
}

type EditProductResp struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
}

type DeleteProductReq struct {
	ID uuid.UUID `json:"id"`
}
