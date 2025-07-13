package models

type Category struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"size:100;not null"`
	Icon   string `gorm:"size:50"`
	TypeID uint
	UserID uint
}