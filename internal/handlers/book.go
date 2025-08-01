package handlers

import (
	"net/http"
	"strconv"

	"go-book-api/internal/services"

	"github.com/gin-gonic/gin"
)

// BookHandler handles book-related requests
type BookHandler struct {
	bookService services.BookService
}

// NewBookHandler creates a new book handler
func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

// GetBooks handles getting all books
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]interface{}
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  books,
		"count": len(books),
	})
}

// GetBook handles getting a single book by ID
// @Summary Get book by ID
// @Description Retrieve a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.bookService.GetBookByID(c.Request.Context(), uint(bookID))
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// CreateBook handles creating a new book
// @Summary Create a new book
// @Description Create a new book with the provided data
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param book body services.CreateBookRequest true "Book data"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req services.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	book, err := h.bookService.CreateBook(c.Request.Context(), req, userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Book created successfully",
		"data":    book,
	})
}

// UpdateBook handles updating a book
// @Summary Update a book
// @Description Update an existing book with new data
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Param book body services.UpdateBookRequest true "Book update data"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var req services.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UpdateBook(c.Request.Context(), uint(bookID), req)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
		"data":    book,
	})
}

// DeleteBook handles deleting a book
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = h.bookService.DeleteBook(c.Request.Context(), uint(bookID))
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
} 