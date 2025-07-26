package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary Create a new book
// @Description สร้างหนังสือใหม่
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.BookCreateRequest true "Book JSON"
// @Success 201 {object} models.JsonResponse
// @Router /api/books [post]
func CreateBook(c *gin.Context) {
	var req models.BookCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{Success: false, Message: "Invalid request data."})
		return
	}

	book := models.Book{Title: req.Title}
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to create book."})
		return
	}

	c.JSON(http.StatusCreated, models.JsonResponse{Success: true, Message: "Book created successfully.", Data: book})
}

// GetBook godoc
// @Summary Get a book by ID
// @Description ดึงข้อมูลหนังสือตาม ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Preload("Users").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{Success: false, Message: "Book not found."})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{Success: true, Message: "Book retrieved successfully.", Data: book})
}

// AddUserToBook godoc
// @Summary Add a user to a book
// @Description เพิ่มผู้ใช้เข้าเล่มหนังสือ
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Param userId path int true "User ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/books/{id}/users/{userId} [post]
func AddUserToBook(c *gin.Context) {
	bookID := c.Param("id")
	userID := c.Param("userId")

	var book models.Book
	if err := database.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{Success: false, Message: "Book not found."})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{Success: false, Message: "User not found."})
		return
	}

	if err := database.DB.Model(&book).Association("Users").Append(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to add user to book."})
		return
	}

	if err := database.DB.Preload("Users").First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to retrieve book."})
		return
	}

	c.JSON(http.StatusOK, models.JsonResponse{Success: true, Message: "User added to book.", Data: book})
}
