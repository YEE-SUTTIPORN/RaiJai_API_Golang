package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTypes godoc
// @Summary Get all types
// @Description ดึงข้อมูลประเภททั้งหมด
// @Tags types
// @Produce json
// @Success 200 {array} models.JsonResponse
// @Router /api/types [get]
func GetTypes(c *gin.Context) {
	var types []models.Type
	database.DB.Find(&types)
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Types retrieved successfully.",
		Data:    types,
	})
}

// CreateType godoc
// @Summary Create a new type
// @Description เพิ่มประเภทใหม่เข้าระบบ
// @Tags types
// @Accept json
// @Produce json
// @Param type body models.TypeCreateRequest true "Type JSON"
// @Success 201 {object} models.JsonResponse
// @Router /api/types [post]
func CreateType(c *gin.Context) {
	var typeRequest models.TypeCreateRequest
	if err := c.ShouldBindJSON(&typeRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{
			Success: false,
			Message: "Invalid input data.",
			Data:    nil,
		})
		return
	}
	var newType models.Type
	newType.Name = typeRequest.Name
	database.DB.Create(&newType)
	c.JSON(http.StatusCreated, models.JsonResponse{
		Success: true,
		Message: "Type created successfully.",
		Data:    newType,
	})
}

// UpdateType godoc
// @Summary Update an existing type
// @Description แก้ไขประเภทที่มีอยู่
// @Tags types
// @Accept json
// @Produce json
// @Param id path int true "Type ID"
// @Param type body models.TypeCreateRequest true "Type JSON"
// @Success 200 {object} models.JsonResponse
// @Router /api/types/{id} [put]
func UpdateType(c *gin.Context) {
	id := c.Param("id")
	var typeRequest models.TypeCreateRequest
	if err := c.ShouldBindJSON(&typeRequest); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{
			Success: false,
			Message: "Invalid input data.",
			Data:    nil,
		})
		return
	}
	var existingType models.Type
	if err := database.DB.First(&existingType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Type not found.",
			Data:    nil,
		})
		return
	}
	existingType.Name = typeRequest.Name
	if err := database.DB.Save(&existingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to update type.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Type updated successfully.",
		Data:    existingType,
	})
}

// DeleteType godoc
// @Summary Delete a type
// @Description ลบประเภทที่มีอยู่
// @Tags types
// @Param id path int true "Type ID"
// @Success 204 {object} models.JsonResponse
// @Router /api/types/{id} [delete]
func DeleteType(c *gin.Context) {
	id := c.Param("id")
	var existingType models.Type
	if err := database.DB.First(&existingType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Type not found.",
			Data:    nil,
		})
		return
	}

	// Delete the type
	if err := database.DB.Delete(&existingType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to delete type.",
			Data:    nil,
		})
		return
	}

	// Return 204 No Content
	c.JSON(http.StatusNoContent, models.JsonResponse{
		Success: true,
		Message: "Type deleted successfully.",
		Data:    nil,
	})
}

// GetTypeByID godoc
// @Summary Get a type by ID
// @Description ดึงข้อมูลประเภทตาม ID
// @Tags types
// @Param id path int true "Type ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/types/{id} [get]
func GetTypeByID(c *gin.Context) {
	id := c.Param("id")
	var typeData models.Type
	if err := database.DB.First(&typeData, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Type not found.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Type retrieved successfully.",
		Data:    typeData,
	})
}