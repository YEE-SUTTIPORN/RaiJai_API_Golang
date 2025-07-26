package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"RaiJaiAPI_Golang/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
