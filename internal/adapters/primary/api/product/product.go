// filepath: /home/frostnzx/dev/go-ecommerce-api/internal/adapters/primary/api/product/product.go
package product

import (
	"encoding/json"
	"net/http"

	coreproduct "github.com/frostnzx/go-ecommerce-api/internal/core/services/product"
	"github.com/google/uuid"
)

type Handler struct {
	svc             coreproduct.API
	authMiddleware  func(http.Handler) http.Handler
	adminMiddleware func(http.Handler) http.Handler
}

func New(svc coreproduct.API, authMiddleware, adminMiddleware func(http.Handler) http.Handler) *Handler {
	return &Handler{
		svc:             svc,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	// Public routes (anyone can view products)
	mux.HandleFunc("GET /products", h.ListProductsHandler)
	mux.HandleFunc("GET /products/{id}", h.GetProductHandler)

	// Admin routes (only admin can manage products)
	mux.Handle("POST /admin/products", h.adminMiddleware(http.HandlerFunc(h.AddProductHandler)))
	mux.Handle("PUT /admin/products/{id}", h.adminMiddleware(http.HandlerFunc(h.EditProductHandler)))
	mux.Handle("DELETE /admin/products/{id}", h.adminMiddleware(http.HandlerFunc(h.DeleteProductHandler)))
}

// DTOs
type addProductReq struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
}

type addProductResp struct {
	ID        uuid.UUID `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	CreatedAt string    `json:"created_at"`
}

type productInfoResp struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
}

type getProductResp struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
	CreatedAt   string    `json:"created_at"`
}

type listProductsResp struct {
	Products []productInfoResp `json:"products"`
}

type editProductReq struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Active      bool    `json:"active"`
}

type editProductResp struct {
	ID          uuid.UUID `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	StockQty    int       `json:"stock_qty"`
	Active      bool      `json:"active"`
}

// Handlers

func (h *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var req addProductReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	in := coreproduct.AddProductReq{
		SKU:         req.SKU,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		StockQty:    req.StockQty,
	}

	res, err := h.svc.AddProduct(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := addProductResp{
		ID:        res.ID,
		SKU:       res.SKU,
		Name:      res.Name,
		CreatedAt: res.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	in := coreproduct.GetProductReq{
		ID: id,
	}

	res, err := h.svc.GetProduct(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := getProductResp{
		ID:          res.ID,
		SKU:         res.SKU,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		StockQty:    res.StockQty,
		Active:      res.Active,
		CreatedAt:   res.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.ListProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var products []productInfoResp
	for _, p := range res.Products {
		products = append(products, productInfoResp{
			ID:          p.ID,
			SKU:         p.SKU,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			StockQty:    p.StockQty,
			Active:      p.Active,
		})
	}

	resp := listProductsResp{Products: products}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) EditProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	var req editProductReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	in := coreproduct.EditProductReq{
		ID:          id,
		SKU:         req.SKU,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		StockQty:    req.StockQty,
		Active:      req.Active,
	}

	res, err := h.svc.EditProduct(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := editProductResp{
		ID:          res.ID,
		SKU:         res.SKU,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		StockQty:    res.StockQty,
		Active:      res.Active,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	in := coreproduct.DeleteProductReq{
		ID: id,
	}

	err = h.svc.DeleteProduct(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
