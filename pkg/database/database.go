package database

import (
	"log"

	"go-book-api/internal/config"
	"go-book-api/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init initializes the database connection and runs migrations
func Init() error {
	dsn := config.GetDatabaseDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	DB = db

	// Run migrations
	if err := runMigrations(); err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// runMigrations runs database migrations
func runMigrations() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Book{},
	)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
