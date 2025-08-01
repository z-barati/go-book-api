# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=go-book-api
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v cmd/main.go

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -cover ./...

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go
	./$(BINARY_NAME)

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	$(GOCMD) fmt ./...

# Vet code
vet:
	$(GOCMD) vet ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Generate swagger docs
swagger:
	swag init -g cmd/main.go -o api/docs

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME)

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

# Database commands
db-migrate:
	$(GOBUILD) -o migrate cmd/migrate/main.go
	./migrate up

db-rollback:
	$(GOBUILD) -o migrate cmd/migrate/main.go
	./migrate down

# Development helpers
dev-setup: deps fmt vet test

# Full build and test
all: clean deps fmt vet test build

.PHONY: build build-linux clean test test-coverage run deps fmt vet lint swagger docker-build docker-run docker-compose-up docker-compose-down db-migrate db-rollback dev-setup all 