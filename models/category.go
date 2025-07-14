package models

type Category struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	Icon   string `gorm:"size:50"`
	TypeID uint
	UserID uint
}

type CategoryCreateRequest struct {
	Name   string `json:"name" binding:"required"`
	Icon   string `json:"icon"`
	TypeID uint   `json:"type_id"`
	UserID uint   `json:"user_id"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Icon string `json:"icon"`
}