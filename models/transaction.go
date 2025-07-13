package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
    ID         uint           `gorm:"primaryKey"`
    Amount     float64        `gorm:"not null"`
    Note       string         `gorm:"size:255"`
    Date       time.Time      `gorm:"not null"`
    UserID     uint
    CategoryID uint
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
}
