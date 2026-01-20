package api

import (
	"net/http"

	"github.com/frostnzx/go-ecommerce-api/internal/core/services/address"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/items"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/order"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/product"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/session"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/user"

	userhandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/user"
)

type App struct {
	server     *http.Server
	userAPI    user.API
	orderAPI   order.API
	addressAPI address.API
	productAPI product.API
	itemsAPI   items.API
	sessionAPI session.API
	port       int
}

func NewApp(userAPI user.API, orderAPI order.API, addressAPI address.API, sessionAPI session.API, addr string) *App {
	mux := http.NewServeMux()

	// compose handlers with core services
	uHandler := userhandler.New(userAPI)
	uHandler.RegisterRoutes(mux)

	// ... register other handlers similarly ...

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &App{
		server:     srv,
		userAPI:    userAPI,
		orderAPI:   orderAPI,
		addressAPI: addressAPI,
	}
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
