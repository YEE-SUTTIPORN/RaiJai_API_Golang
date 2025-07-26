package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	Amount     float64   `gorm:"not null"`
	Note       string    `gorm:"size:255"`
	Date       time.Time `gorm:"not null"`
	UserID     uint
	BookID     uint
	CategoryID uint
	Book       Book
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type TransactionCreateRequest struct {
	Amount     float64   `json:"amount" binding:"required"`
	Note       string    `json:"note"`
	Date       time.Time `json:"date" binding:"required"`
	UserID     uint      `json:"user_id" binding:"required"`
	BookID     uint      `json:"book_id" binding:"required"`
	CategoryID uint      `json:"category_id" binding:"required"`
}

type TransactionUpdateRequest struct {
	Amount     float64   `json:"amount" binding:"required"`
	Note       string    `json:"note"`
	Date       time.Time `json:"date" binding:"required"`
	UserID     uint      `json:"user_id" binding:"required"`
	BookID     uint      `json:"book_id" binding:"required"`
	CategoryID uint      `json:"category_id" binding:"required"`
}
