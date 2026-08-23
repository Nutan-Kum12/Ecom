# Ecom API

A Go REST API for an e-commerce application. It provides user authentication, product management, cart checkout, order persistence, and MySQL migrations.

## Stack

- Go 1.25.1
- Gorilla Mux
- MySQL
- GORM and `database/sql`
- JWT authentication
- `golang-migrate`

## Requirements

- Go 1.25.1 or newer
- MySQL 8 or compatible MySQL server
- Make (optional, for the commands below)

## Getting Started

Clone the repository and install dependencies:

```bash
git clone https://github.com/Nutan-Kum12/Ecom.git
cd Ecom
go mod download
```

Create a local `.env` file. Do not commit it:

```dotenv
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your-password
DB_NAME=ecom
PUBLIC_HOST=http://localhost
PORT=8080
JWT_SECRET=replace-with-a-long-random-secret
JWT_EXPIRATION_IN_SECONDS=604800
```

Create the database, then apply migrations:

```sql
CREATE DATABASE ecom;
```

```bash
make migrate-up
make run
```

The API is available at `http://localhost:8080`.

## API

| Method | Endpoint | Description | Auth |
| --- | --- | --- | --- |
| `GET` | `/health` | Health check | No |
| `POST` | `/api/v1/register` | Register a user | No |
| `POST` | `/api/v1/login` | Get a JWT token | No |
| `GET` | `/api/v1/products` | List products | No |
| `GET` | `/api/v1/products/{id}` | Get a product | No |
| `POST` | `/api/v1/products` | Create a product | Bearer token |
| `POST` | `/api/v1/cart/checkout` | Checkout the cart | Bearer token |

See [api-tests.http](api-tests.http) for REST Client requests and examples.

## Development Commands

```bash
make run             # Start the API
make test            # Run all Go tests
make build           # Build bin/ecom
make migrate-up     # Apply pending migrations
make migrate-down   # Roll back migrations
make db-status      # Check database status
```

Migration files are in `cmd/migrate/migrations/`. Each schema change should include both an `.up.sql` and a `.down.sql` file.

## Project Layout

```text
cmd/             Application, API, database migration, and status entry points
config/          Environment-based configuration
db/              MySQL connection setup
services/        Authentication, users, products, cart, and orders
types/           Shared data models and interfaces
utils/           JSON, validation, and HTTP response helpers
```
