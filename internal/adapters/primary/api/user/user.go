package user

import (
	"encoding/json"
	"net/http"

	coreuser "github.com/frostnzx/go-ecommerce-api/internal/core/services/user"
)

type Handler struct {
	svc coreuser.API
}

func New(svc coreuser.API) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// register concrete paths; method checks happen inside handlers
	mux.HandleFunc("/users", h.RegisterUserHandler)
	// add other routes: /users/login, /users/{id} ...
}

func (h *Handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.RegisterUserReq{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		IsAdmin:  req.IsAdmin,
	}
	res, err := h.svc.RegisterUser(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
