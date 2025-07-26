package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:100;not null"`
	Users     []User `gorm:"many2many:book_users"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type BookCreateRequest struct {
	Title string `json:"title" binding:"required"`
}
