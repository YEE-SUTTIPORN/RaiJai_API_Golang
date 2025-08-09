package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/utils"
)

func TestCreateAndListTransactions(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")
	r := setupTestRouter(t)

	user := models.User{Name: "Bob", Email: "bob@example.com", Password: "pass"}
	if err := database.DB.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}
	ttype := models.Type{Name: "expense"}
	database.DB.Create(&ttype)
	category := models.Category{Name: "food", UserID: user.ID}
	database.DB.Create(&category)
	book := models.Book{Title: "wallet"}
	database.DB.Create(&book)

	token, _ := utils.GenerateToken(user.ID)

	payload := fmt.Sprintf(`{"amount":100,"note":"lunch","date":"%s","user_id":%d,"book_id":%d,"category_id":%d}`,
		time.Now().Format(time.RFC3339), user.ID, book.ID, category.ID)
	req := httptest.NewRequest("POST", "/api/transactions", bytes.NewBufferString(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", w.Code)
	}

	req = httptest.NewRequest("GET", "/api/transactions", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
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
	data, ok := body["data"].([]interface{})
	if !ok || len(data) != 1 {
		t.Fatalf("expected 1 transaction got %v", body["data"])
	}
}
