package database

import (
	"RaiJaiAPI_Golang/config"
	"RaiJaiAPI_Golang/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := config.GetDSN()
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.User{}, &models.Type{}, &models.Category{}, &models.Transaction{})
    DB = db
}