package container

import (
	"go-book-api/internal/repositories"
	"go-book-api/internal/services"
	"go-book-api/pkg/database"

	"gorm.io/gorm"
)

// Container holds all application dependencies
type Container struct {
	DB          *gorm.DB
	UserRepo    repositories.UserRepository
	BookRepo    repositories.BookRepository
	UserService services.UserService
	BookService services.BookService
}

// NewContainer creates a new dependency container
func NewContainer() (*Container, error) {
	// Initialize database
	if err := database.Init(); err != nil {
		return nil, err
	}

	db := database.GetDB()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	bookService := services.NewBookService(bookRepo)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		BookRepo:    bookRepo,
		UserService: userService,
		BookService: bookService,
	}, nil
}

// GetUserService returns the user service
func (c *Container) GetUserService() services.UserService {
	return c.UserService
}

// GetBookService returns the book service
func (c *Container) GetBookService() services.BookService {
	return c.BookService
}

// GetDB returns the database instance
func (c *Container) GetDB() *gorm.DB {
	return c.DB
} 