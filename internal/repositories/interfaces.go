package repositories

import (
	"context"

	"go-book-api/internal/models"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, userID uint) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByUsernameOrEmail(ctx context.Context, username, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID uint) error
}

// BookRepository defines the interface for book data access
type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, bookID uint) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, bookID uint) error
	GetByUserID(ctx context.Context, userID uint) ([]models.Book, error)
} 