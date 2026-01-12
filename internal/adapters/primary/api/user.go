package api

import (
	"net/http"
)

func (a *App) RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", a.CreateUserHandler)
	// ...
}

func (app *App) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement user creation logic using app.userAPI
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created"))
}