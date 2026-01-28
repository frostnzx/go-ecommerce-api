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
	ID uuid.UUID `json:"id"`
}

type addressInfoResp struct {
	ID         uuid.UUID `json:"id"`
	Line1      string    `json:"line1"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	IsDefault  bool      `json:"is_default"`
}

type listAddressesResp struct {
	Addresses []addressInfoResp `json:"addresses"`
}

type getDefaultAddressResp struct {
	ID         uuid.UUID `json:"id"`
	Line1      string    `json:"line1"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
}

// AddAddressHandler godoc
// @Summary      Add a new address
// @Description  Create a new address for the authenticated user
// @Tags         Addresses
// @Accept       json
// @Produce      json
// @Param        request body addAddressReq true "Address data"
// @Success      201 {object} addAddressResp
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Security     BearerAuth
// @Router       /addresses [post]
func (h *Handler) AddAddressHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req addAddressReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	in := coreaddress.AddAddressReq{
		UserID:     claims.ID,
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

	resp := addAddressResp{ID: res.ID}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// ListAddressesHandler godoc
// @Summary      List all addresses
// @Description  Get all addresses for the authenticated user
// @Tags         Addresses
// @Accept       json
// @Produce      json
// @Success      200 {object} listAddressesResp
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /addresses [get]
func (h *Handler) ListAddressesHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	in := coreaddress.ListAddressesReq{UserID: claims.ID}
	res, err := h.svc.ListAddresses(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var addresses []addressInfoResp
	for _, a := range res.Addresses {
		addresses = append(addresses, addressInfoResp{
			ID:         a.ID,
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

// DeleteAddressHandler godoc
// @Summary      Delete an address
// @Description  Delete an address by ID
// @Tags         Addresses
// @Accept       json
// @Produce      json
// @Param        id path string true "Address ID"
// @Success      204 {string} string "No Content"
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Security     BearerAuth
// @Router       /addresses/{id} [delete]
func (h *Handler) DeleteAddressHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	addressIDStr := r.PathValue("id")
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		http.Error(w, "invalid address id", http.StatusBadRequest)
		return
	}

	in := coreaddress.DeleteAddressReq{ID: addressID, UserID: claims.ID}
	err = h.svc.DeleteAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SetDefaultAddressHandler godoc
// @Summary      Set default address
// @Description  Set an address as the default address
// @Tags         Addresses
// @Accept       json
// @Produce      json
// @Param        id path string true "Address ID"
// @Success      204 {string} string "No Content"
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Security     BearerAuth
// @Router       /addresses/{id}/default [put]
func (h *Handler) SetDefaultAddressHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	addressIDStr := r.PathValue("id")
	addressID, err := uuid.Parse(addressIDStr)
	if err != nil {
		http.Error(w, "invalid address id", http.StatusBadRequest)
		return
	}

	in := coreaddress.SetDefaultAddressReq{AddressID: addressID, UserID: claims.ID}
	err = h.svc.SetDefaultAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetDefaultAddressHandler godoc
// @Summary      Get default address
// @Description  Get the default address for the authenticated user
// @Tags         Addresses
// @Accept       json
// @Produce      json
// @Success      200 {object} getDefaultAddressResp
// @Failure      401 {string} string "Unauthorized"
// @Failure      404 {string} string "Not found"
// @Security     BearerAuth
// @Router       /addresses/default [get]
func (h *Handler) GetDefaultAddressHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	in := coreaddress.GetDefaultAddressReq{UserID: claims.ID}
	res, err := h.svc.GetDefaultAddress(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := getDefaultAddressResp{
		ID:         res.ID,
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
