package api

import (
	"net/http"

	"github.com/frostnzx/go-ecommerce-api/internal/core/services/address"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/order"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/user"
)

type App struct {
	server     *http.Server
	userAPI    user.API
	orderAPI   order.API
	addressAPI address.API
	// itemsAPI items.API
	port int
}

func NewApp(userAPI user.API, orderAPI order.API, addressAPI address.API, addr string) *App {
	mux := http.NewServeMux()

	app := &App{
		userAPI:    userAPI,
		orderAPI:   orderAPI,
		addressAPI: addressAPI,
	}

	// register routes
	app.RegisterUserRoutes(mux)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	app.server = srv

	return app
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
