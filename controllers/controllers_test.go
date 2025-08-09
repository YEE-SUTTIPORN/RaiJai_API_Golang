package controllers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/routes"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Type{}, &models.Category{}, &models.Transaction{}, &models.Book{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	database.DB = db
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}
