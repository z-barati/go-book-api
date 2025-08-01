package repositories

import (
	"context"
	"errors"

	"go-book-api/internal/models"

	"gorm.io/gorm"
)

// bookRepository implements BookRepository
type bookRepository struct {
	db *gorm.DB
}

// NewBookRepository creates a new book repository
func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

// Create creates a new book
func (r *bookRepository) Create(ctx context.Context, book *models.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

// GetByID retrieves a book by ID
func (r *bookRepository) GetByID(ctx context.Context, bookID uint) (*models.Book, error) {
	var book models.Book
	err := r.db.WithContext(ctx).Preload("User").First(&book, bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// GetAll retrieves all books
func (r *bookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	err := r.db.WithContext(ctx).Preload("User").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Update updates a book
func (r *bookRepository) Update(ctx context.Context, book *models.Book) error {
	return r.db.WithContext(ctx).Save(book).Error
}

// Delete deletes a book
func (r *bookRepository) Delete(ctx context.Context, bookID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Book{}, bookID).Error
}

// GetByUserID retrieves books by user ID
func (r *bookRepository) GetByUserID(ctx context.Context, userID uint) ([]models.Book, error) {
	var books []models.Book
	err := r.db.WithContext(ctx).Preload("User").Where("created_by = ?", userID).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
} 