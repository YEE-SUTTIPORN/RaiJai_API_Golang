package models

type Category struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	Icon   string `gorm:"size:50"`
	UserID uint   `gorm:"not null"`
	BookID uint   `gorm:"not null"`
	Book   Book   `gorm:"foreignKey:BookID"`
}

type CategoryCreateRequest struct {
	Name   string `json:"name" binding:"required"`
	Icon   string `json:"icon"`
	UserID uint   `json:"user_id"`
	BookID uint   `json:"book_id"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon"`
}