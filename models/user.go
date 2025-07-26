package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;unique;not null"`
	Email     string `gorm:"size:100;not null"`
	Password  string `gorm:"not null"`
	Books     []Book `gorm:"many2many:book_users"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=4,max=50"`
}

type UserUpdateRequest struct {
	Email    string `json:"email" binding:"email,max=100"`
	Password string `json:"password" binding:"min=4,max=50"`
}
