# Go E-Commerce API

A RESTful e-commerce API built with Go, following clean architecture principles (hexagonal architecture). This project provides a complete backend for an e-commerce platform with user authentication, product management, orders, and addresses.

## Project Structure

```
├── cmd/web/                    # Application entry point
├── db/migrations/              # Database migrations
├── internal/
│   ├── adapters/
│   │   ├── primary/api/        # HTTP handlers and middleware
│   │   └── secondary/postgres/ # Database repositories
│   ├── core/
│   │   ├── domain/             # Business entities
│   │   ├── services/           # Business logic
│   │   └── utils/              # Utilities (JWT, etc.)
│   └── ports/                  # Interface definitions
```

## Features

- **User Management**: Registration, login, logout, profile management
- **JWT Authentication**: Access and refresh tokens with session management
- **Role-Based Access**: Admin and regular user roles
- **Product Catalog**: CRUD operations for products (admin only for write operations)
- **Order Management**: Place orders, view order history, cancel orders
- **Address Management**: Multiple addresses per user with default selection
- **Order Items**: Track items within orders with price snapshots

## Tech Stack

- **Go 1.21+**
- **PostgreSQL** - Primary database
- **sqlx** - Database toolkit
- **JWT** - Authentication tokens

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 14+
- [golang-migrate](https://github.com/golang-migrate/migrate) (for database migrations)

### Environment Variables

Create a `.env` file or export the following variables:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=ecommerce
SERVER_PORT=8080
JWT_SECRET=your-secret-key-here
```

### Database Setup

1. Create the database:
```bash
createdb ecommerce
```

2. Run migrations:
```bash
migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/ecommerce?sslmode=disable" up
```

### Running the Server

```bash
go run cmd/web/main.go
```

The server will start on `http://localhost:8080` (or your configured port).

## API Endpoints

### Authentication

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/auth/register` | Register new user | No |
| POST | `/auth/login` | Login and get tokens | No |
| POST | `/auth/logout` | Logout (invalidate session) | Yes |
| POST | `/auth/renew` | Renew access token | No |

### Users

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/users/{id}` | Get user profile | Yes |
| PUT | `/users/{id}` | Update user profile | Yes |
| PUT | `/users/{id}/password` | Change password | Yes |
| DELETE | `/users/{id}` | Delete account | Yes |
| GET | `/admin/users` | List all users | Admin |

### Products

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET | `/products` | List all products | No |
| GET | `/products/{id}` | Get product details | No |
| POST | `/admin/products` | Create product | Admin |
| PUT | `/admin/products/{id}` | Update product | Admin |
| DELETE | `/admin/products/{id}` | Delete product | Admin |

### Orders

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/orders` | Place new order | Yes |
| GET | `/orders` | List user's orders | Yes |
| GET | `/orders/{id}` | Get order details | Yes |
| POST | `/orders/{id}/cancel` | Cancel order | Yes |
| PUT | `/admin/orders/{id}/status` | Update order status | Admin |

### Order Items

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/orders/{orderId}/items` | Add item to order | Yes |
| GET | `/orders/{orderId}/items` | List items in order | Yes |
| GET | `/orders/{orderId}/items/{id}` | Get item details | Yes |
| DELETE | `/orders/{orderId}/items/{id}` | Remove item from order | Yes |
| GET | `/items` | List all user's items | Yes |

### Addresses

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/addresses` | Add new address | Yes |
| GET | `/addresses` | List user's addresses | Yes |
| DELETE | `/addresses/{id}` | Delete address | Yes |
| PUT | `/addresses/{id}/default` | Set as default address | Yes |
| GET | `/addresses/default` | Get default address | Yes |

## Authentication

The API uses JWT tokens for authentication. Include the access token in the `Authorization` header:

```
Authorization: Bearer <access_token>
```

### Token Flow

1. Register or login to receive access and refresh tokens
2. Use access token for authenticated requests
3. When access token expires, use refresh token to get a new one
4. Logout invalidates the session (both tokens become invalid)

## Order Statuses

- `pending` - Order placed, awaiting payment
- `paid` - Payment confirmed
- `shipped` - Order shipped
- `cancelled` - Order cancelled

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o bin/api cmd/web/main.go
```

## Architecture

This project follows hexagonal (ports and adapters) architecture:

- **Domain**: Core business entities and rules
- **Ports**: Interfaces defining how the application interacts with the outside world
- **Adapters**: Implementations of ports (HTTP handlers, database repositories)
- **Services**: Business logic orchestration

This separation allows for easy testing and swapping of implementations (e.g., switching databases).

## License

MIT
