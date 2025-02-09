# Go Windsurf Project

A REST API built with Go, following Domain-Driven Design (DDD) principles.

## Project Structure

The project follows a clean DDD architecture with the following layers:

```
├── cmd
│   └── main.go                 # Application entry point
├── internal
│   ├── domain                  # Domain layer
│   │   ├── models             # Domain models
│   │   └── repositories       # Repository interfaces
│   ├── application            # Application layer
│   │   └── services          # Application services
│   ├── infrastructure         # Infrastructure layer
│   │   ├── persistence       # Database implementations
│   │   └── middleware        # HTTP middleware
│   └── interfaces             # Interface layer
│       └── handlers          # HTTP handlers
├── .env                        # Environment variables
├── .env.example               # Example environment variables
├── go.mod                     # Go module file
├── go.sum                     # Go module checksum
└── README.md                  # This file
```

## Prerequisites

- Go 1.22 or later
- PostgreSQL 12 or later

## Environment Variables

Copy `.env.example` to `.env` and update the values:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=windsurf
DB_PORT=5432
```

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/prakoso-id/go-windsurf.git
   cd go-windsurf
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up the database:
   - Create a PostgreSQL database
   - Update the `.env` file with your database credentials

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`.

## API Endpoints

### Products

- `POST /api/products` - Create a new product
- `GET /api/products/:id` - Get a product by ID
- `PUT /api/products/:id` - Update a product
- `DELETE /api/products/:id` - Delete a product
- `GET /api/products` - List all products

## Architecture

This project follows Domain-Driven Design (DDD) principles with a clean architecture:

1. **Domain Layer** (`internal/domain`)
   - Contains the core business logic
   - Defines domain models and repository interfaces
   - Independent of other layers

2. **Application Layer** (`internal/application`)
   - Implements use cases using domain layer
   - Coordinates between different domains
   - Contains application services

3. **Infrastructure Layer** (`internal/infrastructure`)
   - Implements repository interfaces
   - Handles database connections
   - Provides middleware implementations

4. **Interface Layer** (`internal/interfaces`)
   - Contains HTTP handlers
   - Handles request/response
   - Uses application services

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io) - ORM library
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader
- [golang-jwt](https://github.com/golang-jwt/jwt) - JWT implementation

## Contributing
Contributions are welcome! Please fork the repository and open a pull request with your changes.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
