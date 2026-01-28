package main

import (
	"fmt"
	"log"
	"os"

	"github.com/frostnzx/go-ecommerce-api/internal/adapters/primary/api"
	"github.com/frostnzx/go-ecommerce-api/internal/adapters/secondary/postgres"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/address"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/items"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/order"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/product"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/session"
	"github.com/frostnzx/go-ecommerce-api/internal/core/services/user"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver

	_ "github.com/frostnzx/go-ecommerce-api/docs" // Swagger docs
)

// @title           E-Commerce API
// @version         1.0
// @description     A RESTful e-commerce API built with Go using hexagonal architecture.
// @description     Features: User authentication (JWT), Product management, Order processing, Address management.

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "ecommerce")
	serverPort := getEnv("SERVER_PORT", "8080")

	// Build database connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Connect to database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Ping to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	log.Println("Connected to database successfully")

	// Initialize repositories (secondary adapters)
	userRepo, err := postgres.NewUserRepo(db)
	if err != nil {
		log.Fatalf("failed to create user repository: %v", err)
	}

	sessionRepo, err := postgres.NewSessionRepo(db)
	if err != nil {
		log.Fatalf("failed to create session repository: %v", err)
	}

	addressRepo, err := postgres.NewAddressRepo(db)
	if err != nil {
		log.Fatalf("failed to create address repository: %v", err)
	}

	orderRepo, err := postgres.NewOrderRepo(db)
	if err != nil {
		log.Fatalf("failed to create order repository: %v", err)
	}

	itemsRepo, err := postgres.NewItemsRepo(db)
	if err != nil {
		log.Fatalf("failed to create items repository: %v", err)
	}

	productRepo, err := postgres.NewProductRepo(db)
	if err != nil {
		log.Fatalf("failed to create product repository: %v", err)
	}

	// Initialize services (core business logic)
	sessionService := session.NewService(sessionRepo)
	userService := user.NewService(userRepo, sessionService)
	addressService := address.NewService(addressRepo)
	orderService := order.NewService(orderRepo, itemsRepo, productRepo)
	productService := product.NewService(productRepo)
	itemsService := items.NewService(itemsRepo, productRepo, orderRepo)

	// Create and run HTTP server
	addr := ":" + serverPort
	app := api.NewApp(userService, orderService, addressService, productService, itemsService, addr)

	log.Printf("Starting server on %s", addr)
	if err := app.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// getEnv returns the value of an environment variable or a default value if not set
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
