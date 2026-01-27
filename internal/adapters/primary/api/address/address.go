// filepath: /home/frostnzx/dev/go-ecommerce-api/internal/adapters/primary/api/address/address.go
package address

import (
	"encoding/json"
	"net/http"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/auth"
	coreaddress "github.com/frostnzx/go-ecommerce-api/internal/core/services/address"
	"github.com/google/uuid"
)

type Handler struct {
	svc            coreaddress.API
	authMiddleware func(http.Handler) http.Handler
}

func New(svc coreaddress.API, authMiddleware func(http.Handler) http.Handler) *Handler {
	return &Handler{
		svc:            svc,
		authMiddleware: authMiddleware,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	// All address routes require authentication
	mux.Handle("POST /addresses", h.authMiddleware(http.HandlerFunc(h.AddAddressHandler)))
	mux.Handle("GET /addresses", h.authMiddleware(http.HandlerFunc(h.ListAddressesHandler)))
	mux.Handle("DELETE /addresses/{id}", h.authMiddleware(http.HandlerFunc(h.DeleteAddressHandler)))
	mux.Handle("PUT /addresses/{id}/default", h.authMiddleware(http.HandlerFunc(h.SetDefaultAddressHandler)))
	mux.Handle("GET /addresses/default", h.authMiddleware(http.HandlerFunc(h.GetDefaultAddressHandler)))
}

// DTOs
type addAddressReq struct {
	Line1      string `json:"line1"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type addAddressResp struct {
	ID string `json:"id"`
}

type addressInfoResp struct {
	ID         string `json:"id"`
	Line1      string `json:"line1"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	IsDefault  bool   `json:"is_default"`
}

type listAddressesResp struct {
	Addresses []addressInfoResp `json:"addresses"`
}

type getDefaultAddressResp struct {
	ID         string `json:"id"`
	Line1      string `json:"line1"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

// Handlers

func (h *Handler) AddAddressHandler(w http.ResponseWriter, r *http.Request) {
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

	var req addAddressReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	in := coreaddress.AddAddressReq{
		UserID:     userID,
		Line1:      req.Line1,
		City:       req.City,
		Province:   req.Province,
		PostalCode: req.PostalCode,
		Country:    req.Country,
	}

	res, err := h.svc.AddAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := addAddressResp{
		ID: res.ID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) ListAddressesHandler(w http.ResponseWriter, r *http.Request) {
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

	in := coreaddress.ListAddressesReq{
		UserID: userID,
	}

	res, err := h.svc.ListAddresses(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var addresses []addressInfoResp
	for _, a := range res.Addresses {
		addresses = append(addresses, addressInfoResp{
			ID:         a.ID.String(),
			Line1:      a.Line1,
			City:       a.City,
			Province:   a.Province,
			PostalCode: a.PostalCode,
			Country:    a.Country,
			IsDefault:  a.IsDefault,
		})
	}

	resp := listAddressesResp{Addresses: addresses}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {
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

	addressIDStr := r.PathValue("id")
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		http.Error(w, "invalid address id", http.StatusBadRequest)
		return
	}

	in := coreaddress.DeleteAddressReq{
		ID:     addressID,
		UserID: userID,
	}

	err = h.svc.DeleteAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) SetDefaultAddressHandler(w http.ResponseWriter, r *http.Request) {
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

	addressIDStr := r.PathValue("id")
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		http.Error(w, "invalid address id", http.StatusBadRequest)
		return
	}

	in := coreaddress.SetDefaultAddressReq{
		AddressID: addressID,
		UserID:    userID,
	}

	err = h.svc.SetDefaultAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetDefaultAddressHandler(w http.ResponseWriter, r *http.Request) {
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

	in := coreaddress.GetDefaultAddressReq{
		UserID: userID,
	}

	res, err := h.svc.GetDefaultAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := getDefaultAddressResp{
		ID:         res.ID.String(),
		Line1:      res.Line1,
		City:       res.City,
		Province:   res.Province,
		PostalCode: res.PostalCode,
		Country:    res.Country,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
