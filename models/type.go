package models

type Type struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50;unique;not null"` // "income" or "expense"
}

type TypeCreateRequest struct {
	Name string `json:"name" binding:"required,max=50"`
}