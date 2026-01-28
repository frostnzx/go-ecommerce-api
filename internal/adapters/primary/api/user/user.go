package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/auth"
	coreuser "github.com/frostnzx/go-ecommerce-api/internal/core/services/user"
	"github.com/google/uuid"
)

type Handler struct {
	svc             coreuser.API
	authMiddleware  func(http.Handler) http.Handler
	adminMiddleware func(http.Handler) http.Handler
}

func New(svc coreuser.API, authMiddleware, adminMiddleware func(http.Handler) http.Handler) *Handler {
	return &Handler{
		svc:             svc,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

type registerUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}
type registerUserResp struct {
	ID uuid.UUID `json:"id"`
}

type loginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginUserResp struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Name                  string    `json:"name"`
	Email                 string    `json:"email"`
	IsAdmin               bool      `json:"is_admin"`
}

type logoutUserReq struct {
	SessionID string `json:"session_id"`
}

type renewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
type renewAccessTokenResp struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type updateUserProfileReq struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type getUserProfileReq struct {
	ID uuid.UUID `json:"id"`
}

type getUserProfileResp struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	IsAdmin bool      `json:"is_admin"`
}

type deleteAccountReq struct {
	ID        uuid.UUID `json:"id"`
	SessionID string    `json:"session_id"`
}

type changePasswordProfileReq struct {
	ID              uuid.UUID `json:"id"`
	CurrentPassword string    `json:"current_password"`
	NewPassword     string    `json:"new_password"`
}

// Admin DTOs
type listUsersResp struct {
	Users []userInfo `json:"users"`
}

type userInfo struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	IsAdmin bool      `json:"is_admin"`
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	// Public routes (no auth required)
	mux.HandleFunc("POST /auth/register", h.RegisterUserHandler)
	mux.HandleFunc("POST /auth/login", h.LoginHandler)
	mux.HandleFunc("POST /auth/renew", h.RenewAccessTokenHandler)

	// Protected routes (auth required)
	mux.Handle("GET /users/{id}", h.authMiddleware(http.HandlerFunc(h.GetUserProfileHandler)))
	mux.Handle("PUT /users/{id}", h.authMiddleware(http.HandlerFunc(h.UpdateUserProfileHandler)))
	mux.Handle("PUT /users/{id}/password", h.authMiddleware(http.HandlerFunc(h.ChangePasswordHandler)))
	mux.Handle("DELETE /users/{id}", h.authMiddleware(http.HandlerFunc(h.DeleteAccountHandler)))
	mux.Handle("POST /auth/logout", h.authMiddleware(http.HandlerFunc(h.LogoutHandler)))

	// Admin routes (admin only)
	mux.Handle("GET /admin/users", h.adminMiddleware(http.HandlerFunc(h.ListAllUsersHandler)))
}

// RegisterUserHandler godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body registerUserRequest true "User registration data"
// @Success      201 {object} registerUserResp
// @Failure      400 {string} string "Invalid request"
// @Failure      500 {string} string "Internal server error"
// @Router       /auth/register [post]
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
	resp := registerUserResp{
		ID: res.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(resp)
}

// GetUserProfileHandler godoc
// @Summary      Get user profile
// @Description  Get user profile by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        request body getUserProfileReq true "User ID"
// @Success      200 {object} getUserProfileResp
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /users/{id} [get]
func (h *Handler) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	var req getUserProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.GetUserProfileReq{
		ID: req.ID,
	}
	res, err := h.svc.GetUserProfile(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := getUserProfileResp{
		ID:      res.ID,
		Name:    res.Name,
		Email:   res.Email,
		IsAdmin: res.IsAdmin,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(resp)
}

// UpdateUserProfileHandler godoc
// @Summary      Update user profile
// @Description  Update user profile information
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        request body updateUserProfileReq true "Updated profile data"
// @Success      200 {string} string "OK"
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /users/{id} [put]
func (h *Handler) UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	var req updateUserProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.UpdateUserProfileReq{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	}
	err := h.svc.UpdateUserProfile(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
}

// ChangePasswordHandler godoc
// @Summary      Change password
// @Description  Change user password
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        request body changePasswordProfileReq true "Password change data"
// @Success      204 {string} string "No Content"
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /users/{id}/password [put]
func (h *Handler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req changePasswordProfileReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.ChangePasswordProfileReq{
		ID:              req.ID,
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
	}
	err := h.svc.ChangePassword(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// DeleteAccountHandler godoc
// @Summary      Delete account
// @Description  Delete user account
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        request body deleteAccountReq true "Account deletion data"
// @Success      204 {string} string "No Content"
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /users/{id} [delete]
func (h *Handler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req deleteAccountReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.DeleteAccountReq{
		ID:        req.ID,
		SessionID: claims.SessionID,
	}
	err := h.svc.DeleteAccount(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// LoginHandler godoc
// @Summary      Login user
// @Description  Authenticate user and return tokens
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body loginUserReq true "Login credentials"
// @Success      200 {object} loginUserResp
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Router       /auth/login [post]
func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginUserReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.LoginUserReq{
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := h.svc.LoginUser(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	resp := loginUserResp{
		SessionID:             res.SessionID,
		AccessToken:           res.AccessToken,
		RefreshToken:          res.RefreshToken,
		AccessTokenExpiresAt:  res.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: res.RefreshTokenExpiresAt,
		Name:                  res.Name,
		Email:                 res.Email,
		IsAdmin:               res.IsAdmin,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(resp)
}

// LogoutHandler godoc
// @Summary      Logout user
// @Description  Invalidate user session
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      204 {string} string "No Content"
// @Failure      401 {string} string "Unauthorized"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /auth/logout [post]
func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := auth.GetClaimsFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	in := coreuser.LogoutUserReq{
		SessionID: claims.SessionID,
	}
	err := h.svc.LogoutUser(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204
}

// RenewAccessTokenHandler godoc
// @Summary      Renew access token
// @Description  Get a new access token using refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body renewAccessTokenReq true "Refresh token"
// @Success      200 {object} renewAccessTokenResp
// @Failure      400 {string} string "Invalid request"
// @Failure      401 {string} string "Unauthorized"
// @Router       /auth/renew [post]
func (h *Handler) RenewAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req renewAccessTokenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	in := coreuser.RenewAccessTokenReq{
		RefreshToken: req.RefreshToken,
	}
	res, err := h.svc.RenewAccessToken(r.Context(), in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	resp := renewAccessTokenResp{
		AccessToken:          res.AccessToken,
		AccessTokenExpiresAt: res.AccessTokenExpiresAt,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(resp)
}

// ListAllUsersHandler godoc
// @Summary      List all users (Admin)
// @Description  Get a list of all users (admin only)
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Success      200 {object} listUsersResp
// @Failure      401 {string} string "Unauthorized"
// @Failure      403 {string} string "Forbidden"
// @Failure      500 {string} string "Internal server error"
// @Security     BearerAuth
// @Router       /admin/users [get]
func (h *Handler) ListAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []userInfo
	for _, u := range res.Users {
		users = append(users, userInfo{
			ID:      u.ID,
			Name:    u.Name,
			Email:   u.Email,
			IsAdmin: u.IsAdmin,
		})
	}

	resp := listUsersResp{Users: users}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(resp)
}
