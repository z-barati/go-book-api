package services

import (
	"context"

	"go-book-api/internal/models"
)

// UserService defines the interface for user-related business logic
type UserService interface {
	Register(ctx context.Context, req RegisterRequest) (*models.User, string, error)
	Login(ctx context.Context, req LoginRequest) (*models.User, string, error)
	RefreshToken(ctx context.Context, userID uint) (string, error)
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
}

// BookService defines the interface for book-related business logic
type BookService interface {
	GetBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, bookID uint) (*models.Book, error)
	CreateBook(ctx context.Context, req CreateBookRequest, userID uint) (*models.Book, error)
	UpdateBook(ctx context.Context, bookID uint, req UpdateBookRequest) (*models.Book, error)
	DeleteBook(ctx context.Context, bookID uint) error
}

// RegisterRequest represents the registration request
type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateBookRequest represents the book creation request
type CreateBookRequest struct {
	Title       string  `json:"title" binding:"required"`
	Author      string  `json:"author" binding:"required"`
	Description string  `json:"description"`
	ISBN        string  `json:"isbn"`
	Pages       int     `json:"pages"`
	Genre       string  `json:"genre"`
	Language    string  `json:"language"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

// UpdateBookRequest represents the book update request
type UpdateBookRequest struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	ISBN        string  `json:"isbn"`
	Pages       int     `json:"pages"`
	Genre       string  `json:"genre"`
	Language    string  `json:"language"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
} 