# Go MongoDB Auth Service

A production-ready Authentication and Data Service built with Go, MongoDB, and clean architecture principles.

## Features

- **Architecture**: Domain-Driven Design (DDD) with Repository-Service pattern.
- **Authentication**: User registration and login with JWT (JSON Web Tokens).
- **Security**: Password hashing using bcrypt.
- **Database**: Native MongoDB driver implementation.
- **Admin Dashboard**: Aggregated user statistics and data overview.

## Project Structure

```
.
├── cmd/api          # Application entry point
├── internal/
│   ├── config       # Configuration loader
│   ├── database     # Database connections
│   ├── handlers     # HTTP transport layer
│   ├── middleware   # Auth middleware
│   ├── models       # Domain entities
│   ├── repository   # Data access layer
│   └── service      # Business logic
└── pkg/
    └── logger       # Structured logging
```

## Getting Started

### Prerequisites

- Go 1.20+
- MongoDB

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-mongo-auth-service.git
   ```

2. Configuration:
   Set the following environment variables (optional):
   - `MONGO_URI`: MongoDB connection string (default: `mongodb://localhost:27017`)
   - `JWT_SECRET`: Secret key for signing tokens
   - `PORT`: Server port (default: `:8080`)

3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

### Auth
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login and receive JWT

### Data
- `GET /api/v1/data` - Retrieve protected user data (Requires Bearer Token)

### Admin
- `GET /api/v1/admin/overview` - View system-wide user statistics (Requires Bearer Token)
