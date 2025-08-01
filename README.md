# Go Book API

A RESTful API for managing books built with Go, featuring clean architecture, comprehensive testing, and modern development practices.

## 🚀 Features

- **RESTful API**: Full CRUD operations for books
- **Clean Architecture**: Separation of concerns with layers (handlers, services, repositories)
- **Database Integration**: MySQL with GORM ORM
- **Authentication**: JWT-based authentication
- **Validation**: Request validation and error handling
- **Testing**: Unit tests and integration tests
- **Documentation**: Auto-generated Swagger API documentation
- **Docker Support**: Containerized application
- **Configuration**: TOML-based configuration with Viper

## 📋 Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Docker (optional)
- Make (optional, for build scripts)

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/go-book-api.git
   cd go-book-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up configuration**
   ```bash
   cp config.toml-example config.toml
   # Edit config.toml with your database credentials
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## 🏗️ Project Structure

```
go-book-api/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── handlers/               # HTTP handlers
│   ├── services/               # Business logic
│   ├── repositories/           # Data access layer
│   ├── models/                 # Data models
│   ├── middleware/             # HTTP middleware
│   └── config/                 # Configuration
├── pkg/
│   ├── database/               # Database utilities
│   ├── auth/                   # Authentication utilities
│   └── utils/                  # Common utilities
├── api/
│   └── docs/                   # API documentation
├── scripts/                    # Build and deployment scripts
├── tests/                      # Integration tests
├── docker/                     # Docker configuration
├── config.toml-example        # Configuration template
├── go.mod                     # Go module file
├── go.sum                     # Go module checksums
├── Dockerfile                 # Docker configuration
├── docker-compose.yml         # Docker Compose configuration
└── README.md                  # This file
```

## 🚀 Quick Start with Docker

1. **Start the application with Docker Compose**
   ```bash
   docker-compose up -d
   ```

2. **Access the API**
   - API Base URL: `http://localhost:8080`
   - Health Check: `http://localhost:8080/health`
   - Swagger Documentation: `http://localhost:8080/swagger/index.html`

## 📚 API Endpoints

### Books
- `GET /api/v1/books` - List all books
- `GET /api/v1/books/{id}` - Get book by ID
- `POST /api/v1/books` - Create a new book
- `PUT /api/v1/books/{id}` - Update a book
- `DELETE /api/v1/books/{id}` - Delete a book

### Authentication
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/refresh` - Refresh token

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run integration tests:
```bash
go test ./tests/...
```

## 🔧 Development

### Code Quality
- **Linting**: `golangci-lint run`
- **Formatting**: `go fmt ./...`
- **Vet**: `go vet ./...`

### Database Migrations
```bash
# Run migrations
go run cmd/migrate/main.go up

# Rollback migrations
go run cmd/migrate/main.go down
```

### Generate Swagger Documentation
```bash
make swagger
# or
swag init -g cmd/main.go -o api/docs
```

## 📦 Build

Build for production:
```bash
go build -o bin/go-book-api cmd/main.go
```

Build with Docker:
```bash
docker build -t go-book-api .
```

## 🚀 Deployment

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_NAME` | Database name | `book_db` |
| `DB_USER` | Database user | `root` |
| `DB_PASSWORD` | Database password | `password` |
| `JWT_SECRET` | JWT secret key | `your-secret-key` |
| `PORT` | Server port | `8080` |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT implementation
- [Viper](https://github.com/spf13/viper) - Configuration management

## 📞 Support

If you have any questions or need help, please open an issue on GitHub.