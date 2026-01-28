package order

import (
	"encoding/json"
	"net/http"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/auth"
	coreorder "github.com/frostnzx/go-ecommerce-api/internal/core/services/order"
	"github.com/google/uuid"
)

type Handler struct {
	svc             coreorder.API
	authMiddleware  func(http.Handler) http.Handler
	adminMiddleware func(http.Handler) http.Handler
}

func New(svc coreorder.API, authMiddleware, adminMiddleware func(http.Handler) http.Handler) *Handler {
	return &Handler{
		svc:             svc,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	// Protected routes (auth required)
	mux.Handle("POST /orders", h.authMiddleware(http.HandlerFunc(h.PlaceOrderHandler)))
	mux.Handle("GET /orders", h.authMiddleware(http.HandlerFunc(h.ListOrdersHandler)))
	mux.Handle("GET /orders/{id}", h.authMiddleware(http.HandlerFunc(h.GetOrderHandler)))
	mux.Handle("POST /orders/{id}/cancel", h.authMiddleware(http.HandlerFunc(h.CancelOrderHandler)))

	// Admin routes
	mux.Handle("PUT /admin/orders/{id}/status", h.adminMiddleware(http.HandlerFunc(h.UpdateOrderStatusHandler)))
}

// DTOs
type orderItemReq struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type placeOrderReq struct {
	AddressID string         `json:"address_id"`
	Items     []orderItemReq `json:"items"`
}

type placeOrderResp struct {
	ID          uuid.UUID `json:"id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   string    `json:"created_at"`
}

type orderInfoResp struct {
	ID          uuid.UUID `json:"id"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   string    `json:"created_at"`
}

type listOrdersResp struct {
	Orders []orderInfoResp `json:"orders"`
}

type orderItemInfoResp struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
}

type getOrderResp struct {
	ID          uuid.UUID           `json:"id"`
	UserID      uuid.UUID           `json:"user_id"`
	AddressID   uuid.UUID           `json:"address_id"`
	Status      string              `json:"status"`
	TotalAmount float64             `json:"total_amount"`
	Items       []orderItemInfoResp `json:"items"`
	CreatedAt   string              `json:"created_at"`
}

type updateStatusReq struct {
	Status string `json:"status"`
}

// Handlers

func (h *Handler) PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req placeOrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	addressID, err := uuid.Parse(req.AddressID)
	if err != nil {
		http.Error(w, "invalid address id", http.StatusBadRequest)
		return
	}

	var items []coreorder.OrderItemReq
	for _, item := range req.Items {
		productID, err := uuid.Parse(item.ProductID)
		if err != nil {
			http.Error(w, "invalid product id", http.StatusBadRequest)
			return
		}
		items = append(items, coreorder.OrderItemReq{
			ProductID: productID,
			Quantity:  item.Quantity,
		})
	}

	in := coreorder.PlaceOrderReq{
		UserID:    claims.ID,
		AddressID: addressID,
		Items:     items,
	}

	res, err := h.svc.PlaceOrder(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := placeOrderResp{
		ID:          res.ID,
		TotalAmount: res.TotalAmount,
		Status:      res.Status,
		CreatedAt:   res.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	in := coreorder.ListOrdersReq{
		UserID: claims.ID,
	}

	res, err := h.svc.ListOrders(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var orders []orderInfoResp
	for _, o := range res.Orders {
		orders = append(orders, orderInfoResp{
			ID:          o.ID,
			Status:      o.Status,
			TotalAmount: o.TotalAmount,
			CreatedAt:   o.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	resp := listOrdersResp{Orders: orders}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	orderIDStr := r.PathValue("id")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	in := coreorder.GetOrderReq{
		OrderID: orderID,
		UserID:  claims.ID,
	}

	res, err := h.svc.GetOrder(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var items []orderItemInfoResp
	for _, item := range res.Items {
		items = append(items, orderItemInfoResp{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}

	resp := getOrderResp{
		ID:          res.ID,
		UserID:      res.UserID,
		AddressID:   res.AddressID,
		Status:      res.Status,
		TotalAmount: res.TotalAmount,
		Items:       items,
		CreatedAt:   res.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) CancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	orderIDStr := r.PathValue("id")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	in := coreorder.CancelOrderReq{
		OrderID: orderID,
		UserID:  claims.ID,
	}

	err = h.svc.CancelOrder(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.PathValue("id")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	var req updateStatusReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	in := coreorder.UpdateOrderStatusReq{
		OrderID: orderID,
		Status:  req.Status,
	}

	err = h.svc.UpdateOrderStatus(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
