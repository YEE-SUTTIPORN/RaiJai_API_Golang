package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary User login
// @Description User login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login JSON"
// @Success 200 {object} models.JsonResponse
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{Success: false, Message: "Invalid input."})
		return
	}
	var user models.User
	if err := database.DB.Where("name = ?", req.Name).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, models.JsonResponse{Success: false, Message: "Invalid credentials."})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.JsonResponse{Success: false, Message: "Invalid credentials."})
		return
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to generate token."})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{Success: true, Message: "Login successful.", Data: gin.H{"token": token}})
}



// Register godoc
// @Summary Add a new user
// @Description เพิ่มผู้ใช้ใหม่เข้าระบบ
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.UserCreateRequest true "User JSON"
// @Success 201 {object} models.JsonResponse
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
    var req models.UserCreateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.JsonResponse{
            Success: false,
            Message: "Invalid input data.",
            Data:    nil,
        })
        return
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to process password."})
        return
    }

    user := models.User{
        Name:      req.Name,
        Email:     req.Email,
        Password:  string(hashed),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Validate unique username
    var existingUser models.User
    if err := database.DB.Where("name = ?", user.Name).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, models.JsonResponse{
            Success: false,
            Message: "Username already exists.",
            Data:    nil,
        })
        return
    }

    database.DB.Create(&user)
    c.JSON(http.StatusCreated, models.JsonResponse{
        Success: true,
        Message: "User created successfully.",
        Data:    user,
    })
}
