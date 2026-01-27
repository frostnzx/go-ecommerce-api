package api

import (
	"net/http"
	"os"

	"github.com/frostnzx/go-ecommerce-api/internal/core/services/address"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/items"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/order"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/product"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/user"
	"github.com/frostnzx/go-ecommerce-api/internal/core/utils"

	addresshandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/address"
	itemshandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/items"
	orderhandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/order"
	producthandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/product"
	userhandler "github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api/user"
)

type App struct {
	server     *http.Server
	userAPI    user.API
	orderAPI   order.API
	addressAPI address.API
	productAPI product.API
	itemsAPI   items.API
}

func NewApp(userAPI user.API, orderAPI order.API, addressAPI address.API, productAPI product.API, itemsAPI items.API, addr string) *App {
	mux := http.NewServeMux()

	// Create JWT maker and auth middleware
	secretKey := os.Getenv("JWT_SECRET")
	tokenMaker := utils.NewJWTMaker(secretKey)
	authMiddleware := GetAuthMiddlewareFunc(tokenMaker)
	adminMiddleware := GetAdminMiddlewareFunc(tokenMaker)

	// Compose handlers with core services and middleware
	uHandler := userhandler.New(userAPI, authMiddleware, adminMiddleware)
	uHandler.SetupRoutes(mux)

	oHandler := orderhandler.New(orderAPI, authMiddleware, adminMiddleware)
	oHandler.SetupRoutes(mux)

	pHandler := producthandler.New(productAPI, authMiddleware, adminMiddleware)
	pHandler.SetupRoutes(mux)

	aHandler := addresshandler.New(addressAPI, authMiddleware)
	aHandler.SetupRoutes(mux)

	iHandler := itemshandler.New(itemsAPI, authMiddleware)
	iHandler.SetupRoutes(mux)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &App{
		server:     srv,
		userAPI:    userAPI,
		orderAPI:   orderAPI,
		addressAPI: addressAPI,
		productAPI: productAPI,
		itemsAPI:   itemsAPI,
	}
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}
