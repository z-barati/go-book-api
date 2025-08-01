package main

import (
	"log"
	"net/http"

	"go-book-api/api/docs"
	"go-book-api/internal/config"
	"go-book-api/internal/container"
	"go-book-api/internal/handlers"
	"go-book-api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Book API
// @version 1.0
// @description A RESTful API for managing books
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize Swagger docs
	_ = docs.SwaggerInfo

	// Initialize dependency container
	container, err := container.NewContainer()
	if err != nil {
		log.Fatal("Failed to initialize container:", err)
	}

	// Set Gin mode
	if viper.GetString("app.mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create handlers with dependency injection
	authHandler := handlers.NewAuthHandler(container.GetUserService())
	bookHandler := handlers.NewBookHandler(container.GetBookService())

	// Create router
	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Go Book API is running",
		})
	})

	// Test route
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test route working"})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// API routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Book routes
			books := protected.Group("/books")
			{
				books.GET("/", bookHandler.GetBooks)
				books.GET("/:id", bookHandler.GetBook)
				books.POST("/", bookHandler.CreateBook)
				books.PUT("/:id", bookHandler.UpdateBook)
				books.DELETE("/:id", bookHandler.DeleteBook)
			}
		}
	}

	// Start server
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
