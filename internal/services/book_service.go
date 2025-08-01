package services

import (
	"context"
	"errors"

	"go-book-api/internal/models"
	"go-book-api/internal/repositories"
)

// bookService implements BookService
type bookService struct {
	bookRepo repositories.BookRepository
}

// NewBookService creates a new book service
func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

// GetBooks retrieves all books
func (s *bookService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return s.bookRepo.GetAll(ctx)
}

// GetBookByID retrieves a book by ID
func (s *bookService) GetBookByID(ctx context.Context, bookID uint) (*models.Book, error) {
	return s.bookRepo.GetByID(ctx, bookID)
}

// CreateBook creates a new book
func (s *bookService) CreateBook(ctx context.Context, req CreateBookRequest, userID uint) (*models.Book, error) {
	book := &models.Book{
		Title:       req.Title,
		Author:      req.Author,
		Description: req.Description,
		ISBN:        req.ISBN,
		Pages:       req.Pages,
		Genre:       req.Genre,
		Language:    req.Language,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedBy:   userID,
	}

	if err := s.bookRepo.Create(ctx, book); err != nil {
		return nil, errors.New("failed to create book")
	}

	// Reload the book with user data
	createdBook, err := s.bookRepo.GetByID(ctx, book.ID)
	if err != nil {
		return nil, errors.New("failed to retrieve created book")
	}

	return createdBook, nil
}

// UpdateBook updates an existing book
func (s *bookService) UpdateBook(ctx context.Context, bookID uint, req UpdateBookRequest) (*models.Book, error) {
	// Get existing book
	book, err := s.bookRepo.GetByID(ctx, bookID)
	if err != nil {
		return nil, errors.New("book not found")
	}

	// Update fields if provided
	if req.Title != "" {
		book.Title = req.Title
	}
	if req.Author != "" {
		book.Author = req.Author
	}
	if req.Description != "" {
		book.Description = req.Description
	}
	if req.ISBN != "" {
		book.ISBN = req.ISBN
	}
	if req.Pages > 0 {
		book.Pages = req.Pages
	}
	if req.Genre != "" {
		book.Genre = req.Genre
	}
	if req.Language != "" {
		book.Language = req.Language
	}
	if req.Price > 0 {
		book.Price = req.Price
	}
	if req.Stock >= 0 {
		book.Stock = req.Stock
	}

	if err := s.bookRepo.Update(ctx, book); err != nil {
		return nil, errors.New("failed to update book")
	}

	// Reload the book with user data
	updatedBook, err := s.bookRepo.GetByID(ctx, book.ID)
	if err != nil {
		return nil, errors.New("failed to retrieve updated book")
	}

	return updatedBook, nil
}

// DeleteBook deletes a book
func (s *bookService) DeleteBook(ctx context.Context, bookID uint) error {
	// Check if book exists
	_, err := s.bookRepo.GetByID(ctx, bookID)
	if err != nil {
		return errors.New("book not found")
	}

	return s.bookRepo.Delete(ctx, bookID)
} 