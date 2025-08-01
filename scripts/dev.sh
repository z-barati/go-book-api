#!/bin/bash

# Development script for Go Book API

set -e

echo "🚀 Go Book API Development Script"
echo "=================================="

case "$1" in
    "setup")
        echo "📦 Setting up development environment..."
        go mod tidy
        go mod download
        echo "✅ Dependencies installed"
        ;;
    
    "run")
        echo "🏃 Running the application..."
        go run cmd/main.go
        ;;
    
    "test")
        echo "🧪 Running tests..."
        go test -v ./...
        ;;
    
    "test-coverage")
        echo "🧪 Running tests with coverage..."
        go test -cover ./...
        ;;
    
    "build")
        echo "🔨 Building the application..."
        go build -o bin/go-book-api cmd/main.go
        echo "✅ Build complete: bin/go-book-api"
        ;;
    
    "fmt")
        echo "🎨 Formatting code..."
        go fmt ./...
        echo "✅ Code formatted"
        ;;
    
    "vet")
        echo "🔍 Running go vet..."
        go vet ./...
        echo "✅ Code vetted"
        ;;
    
    "lint")
        echo "🔍 Running linter..."
        if command -v golangci-lint &> /dev/null; then
            golangci-lint run
            echo "✅ Linting complete"
        else
            echo "⚠️  golangci-lint not found. Install it first:"
            echo "   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
        fi
        ;;
    
    "docker-build")
        echo "🐳 Building Docker image..."
        docker build -t go-book-api .
        echo "✅ Docker image built"
        ;;
    
    "docker-run")
        echo "🐳 Running with Docker Compose..."
        docker-compose up -d
        echo "✅ Application running on http://localhost:8080"
        ;;
    
    "docker-stop")
        echo "🐳 Stopping Docker containers..."
        docker-compose down
        echo "✅ Containers stopped"
        ;;
    
    "clean")
        echo "🧹 Cleaning build artifacts..."
        go clean
        rm -f bin/go-book-api
        rm -f go-book-api
        echo "✅ Clean complete"
        ;;
    
    "help"|*)
        echo "Usage: $0 {setup|run|test|test-coverage|build|fmt|vet|lint|docker-build|docker-run|docker-stop|clean|help}"
        echo ""
        echo "Commands:"
        echo "  setup         - Install dependencies"
        echo "  run           - Run the application"
        echo "  test          - Run tests"
        echo "  test-coverage - Run tests with coverage"
        echo "  build         - Build the application"
        echo "  fmt           - Format code"
        echo "  vet           - Run go vet"
        echo "  lint          - Run linter"
        echo "  docker-build  - Build Docker image"
        echo "  docker-run    - Run with Docker Compose"
        echo "  docker-stop   - Stop Docker containers"
        echo "  clean         - Clean build artifacts"
        echo "  help          - Show this help"
        ;;
esac 