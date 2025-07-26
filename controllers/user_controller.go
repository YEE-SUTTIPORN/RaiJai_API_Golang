package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers godoc
// @Summary Get all users
// @Description ดึงข้อมูลผู้ใช้ทั้งหมด
// @Tags users
// @Produce json
// @Success 200 {array} models.JsonResponse
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, models.JsonResponse{ Success: true, Message: "Users retrieved successfully.", Data: users })
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description แก้ไขข้อมูลผู้ใช้ที่มีอยู่
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserUpdateRequest true "User JSON"
// @Success 200 {object} models.JsonResponse
// @Router /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    println("UpdateUser ID:", id)

    if id == "" {
        c.JSON(http.StatusBadRequest, 
            models.JsonResponse{
            Success: false,
            Message: "ID is required.",})
        return
    }

    var req models.UserUpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, models.JsonResponse {
            Success: false,
            Message: "Invalid input data.",})
        return
    }

    var existingUser models.User
    if err := database.DB.First(&existingUser, id).Error; err != nil {
        c.JSON(http.StatusNotFound, 
            models.JsonResponse{
            Success: false,
            Message: "User not found.",})
        return
    }
 
    if req.Email != "" {
        existingUser.Email = req.Email
    }
    if req.Password != "" {
        hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
        if err != nil {
            c.JSON(http.StatusInternalServerError, models.JsonResponse{Success: false, Message: "Failed to process password."})
            return
        }
        existingUser.Password = string(hashed)
    }
    existingUser.UpdatedAt = time.Now()

    if err := database.DB.Save(&existingUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.JsonResponse {
            Success: false,
            Message: "Failed to update user.",
        })
        return
    }

    c.JSON(http.StatusOK, models.JsonResponse { 
        Success: true,
        Message: "User updated successfully.",
        Data:    existingUser,
     })
}


// DeleteUser godoc
// @Summary Delete a user
// @Description ลบผู้ใช้ตาม ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, models.JsonResponse {
            Success: false,
            Message: "ID is required.",
        })
        return
    }

    // Check if user exists
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, models.JsonResponse {
            Success: false,
            Message: "User not found.",
        })
        return
    }

    if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, models.JsonResponse {
            Success: false,
            Message: "Failed to delete user.",
        })
        return
    }

    c.JSON(http.StatusOK, models.JsonResponse {
        Success: true,
        Message: "User deleted successfully.",
    })
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description ดึงข้อมูลผู้ใช้ตาม ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} models.JsonResponse
// @Router /api/users/{id} [get]
func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, models.JsonResponse {
            Success: false,
            Message: "ID is required.",
        })
        return
    }

    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, models.JsonResponse {
            Success: false,
            Message: "User not found.",
        })
        return
    }

    c.JSON(http.StatusOK, models.JsonResponse {
        Success: true,
        Message: "User retrieved successfully.",
        Data:    user,
    })
}
