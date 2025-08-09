package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/utils"
)

func TestGetUsers(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	r := setupTestRouter(t)
	u := models.User{Name: "Alice", Email: "alice@example.com", Password: "pass"}
	if err := database.DB.Create(&u).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}
	token, _ := utils.GenerateToken(u.ID)

	req := httptest.NewRequest("GET", "/api/users/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if body["success"] != true {
		t.Fatalf("expected success")
	}
}
