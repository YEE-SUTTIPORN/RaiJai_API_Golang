package controllers_test

import (
	"RaiJaiAPI_Golang/controllers"
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	database.ConnectDB()
	gin.SetMode(gin.TestMode) // Set Gin to test mode to avoid logging during tests

	r := gin.Default()
	r.GET("/api/types", controllers.GetTypes)
	r.GET("/api/types/:id", controllers.GetTypeByID)
	r.POST("/api/types", controllers.CreateType)
	r.PUT("/api/types/:id", controllers.UpdateType)
	r.DELETE("/api/types/:id", controllers.DeleteType)
	return r
}

func TestGetTypes(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/api/types", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateType(t *testing.T) {
	r := setupRouter()
	body := models.TypeCreateRequest{Name: "test-type"}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/api/types", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateType_NotFound(t *testing.T) {
	r := setupRouter()
	body := models.TypeCreateRequest{Name: "updated-type"}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "/api/types/9999", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteType_NotFound(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/types/9999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
