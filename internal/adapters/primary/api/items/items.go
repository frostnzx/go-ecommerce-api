// filepath: /home/frostnzx/dev/go-ecommerce-api/internal/adapters/primary/api/items/items.go
package items

import (
	"encoding/json"
	"net/http"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/auth"
	coreitems "github.com/frostnzx/go-ecommerce-api/internal/core/services/items"
	"github.com/google/uuid"
)

type Handler struct {
	svc            coreitems.API
	authMiddleware func(http.Handler) http.Handler
}

func New(svc coreitems.API, authMiddleware func(http.Handler) http.Handler) *Handler {
	return &Handler{
		svc:            svc,
		authMiddleware: authMiddleware,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	// All item routes require authentication
	mux.Handle("POST /orders/{orderId}/items", h.authMiddleware(http.HandlerFunc(h.AddItemHandler)))
	mux.Handle("GET /orders/{orderId}/items", h.authMiddleware(http.HandlerFunc(h.ListItemsByOrderHandler)))
	mux.Handle("GET /orders/{orderId}/items/{id}", h.authMiddleware(http.HandlerFunc(h.GetItemHandler)))
	mux.Handle("DELETE /orders/{orderId}/items/{id}", h.authMiddleware(http.HandlerFunc(h.DeleteItemHandler)))
	mux.Handle("GET /items", h.authMiddleware(http.HandlerFunc(h.ListItemsByUserHandler)))
}

// DTOs
type addItemReq struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type addItemResp struct {
	ID                string  `json:"id"`
	OrderID           string  `json:"order_id"`
	ProductID         string  `json:"product_id"`
	Quantity          int     `json:"quantity"`
	UnitPriceSnapshot float64 `json:"unit_price_snapshot"`
}

type itemInfoResp struct {
	ID                string  `json:"id"`
	ProductID         string  `json:"product_id"`
	Quantity          int     `json:"quantity"`
	UnitPriceSnapshot float64 `json:"unit_price_snapshot"`
}

type listItemsResp struct {
	Items []itemInfoResp `json:"items"`
}

type getItemResp struct {
	ID                string  `json:"id"`
	OrderID           string  `json:"order_id"`
	ProductID         string  `json:"product_id"`
	Quantity          int     `json:"quantity"`
	UnitPriceSnapshot float64 `json:"unit_price_snapshot"`
}

// Handlers

func (h *Handler) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	orderIDStr := r.PathValue("orderId")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	var req addItemReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	productID, err := uuid.Parse(req.ProductID)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	// Note: AddItem doesn't verify ownership - you may want to add that
	_ = userID // We could add ownership check if needed

	in := coreitems.AddItemReq{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  req.Quantity,
	}

	res, err := h.svc.AddItem(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := addItemResp{
		ID:                res.ID.String(),
		OrderID:           res.OrderID.String(),
		ProductID:         res.ProductID.String(),
		Quantity:          res.Quantity,
		UnitPriceSnapshot: res.UnitPriceSnapshot,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetItemHandler(w http.ResponseWriter, r *http.Request) {
	itemIDStr := r.PathValue("id")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}

	in := coreitems.GetItemReq{
		ID: itemID,
	}

	res, err := h.svc.GetItem(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := getItemResp{
		ID:                res.ID.String(),
		OrderID:           res.OrderID.String(),
		ProductID:         res.ProductID.String(),
		Quantity:          res.Quantity,
		UnitPriceSnapshot: res.UnitPriceSnapshot,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListItemsByOrderHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	orderIDStr := r.PathValue("orderId")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	in := coreitems.ListItemsByOrderReq{
		OrderID: orderID,
		UserID:  userID,
	}

	res, err := h.svc.ListItemsByOrder(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var items []itemInfoResp
	for _, item := range res.Items {
		items = append(items, itemInfoResp{
			ID:                item.ID.String(),
			ProductID:         item.ProductID.String(),
			Quantity:          item.Quantity,
			UnitPriceSnapshot: item.UnitPriceSnapshot,
		})
	}

	resp := listItemsResp{Items: items}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListItemsByUserHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	in := coreitems.ListItemsByUserReq{
		UserID: userID,
	}

	res, err := h.svc.ListItemsByUser(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var items []itemInfoResp
	for _, item := range res.Items {
		items = append(items, itemInfoResp{
			ID:                item.ID.String(),
			ProductID:         item.ProductID.String(),
			Quantity:          item.Quantity,
			UnitPriceSnapshot: item.UnitPriceSnapshot,
		})
	}

	resp := listItemsResp{Items: items}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	orderIDStr := r.PathValue("orderId")
	orderID, err := uuid.Parse(orderIDStr)
	if err != nil {
		http.Error(w, "invalid order id", http.StatusBadRequest)
		return
	}

	itemIDStr := r.PathValue("id")
	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}

	in := coreitems.DeleteItemReq{
		ID:      itemID,
		OrderID: orderID,
		UserID:  userID,
	}

	err = h.svc.DeleteItem(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
