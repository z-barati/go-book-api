package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents a book in the system
type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null;size:200"`
	Author      string         `json:"author" gorm:"not null;size:100"`
	Description string         `json:"description" gorm:"type:text"`
	ISBN        string         `json:"isbn" gorm:"uniqueIndex;size:20"`
	PublishedAt *time.Time     `json:"published_at"`
	Pages       int            `json:"pages"`
	Genre       string         `json:"genre" gorm:"size:50"`
	Language    string         `json:"language" gorm:"size:20;default:'en'"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	Stock       int            `json:"stock" gorm:"default:0"`
	CreatedBy   uint           `json:"created_by" gorm:"not null"`
	User        User           `json:"user" gorm:"foreignKey:CreatedBy"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName specifies the table name for Book
func (Book) TableName() string {
	return "books"
} 