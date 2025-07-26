package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description เพิ่มหมวดหมู่
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.CategoryCreateRequest true "Category JSON"
// @Success 201 {object} models.JsonResponse
// @Router /api/categories [post]
func CreateCategory(c *gin.Context) {
	var request models.CategoryCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse { 
			Success: false,
			Message: "Invalid request data.",
			Data:    nil,
		})
		return
	}

	isExists := database.DB.Where("name = ?", request.Name).First(&models.Category{}).RowsAffected > 0
	if isExists {
		c.JSON(http.StatusConflict, models.JsonResponse{
			Success: false,
			Message: "Category with this name already exists.",
			Data:    nil,
		})
		return
	}

	category := models.Category{
		Name:   request.Name,
		Icon:   request.Icon,
		TypeID: request.TypeID,
		UserID: request.UserID,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse {
			Success: false,
			Message: "Failed to create category.",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.JsonResponse{
		Success: true,
		Message: "Category created successfully.",
		Data:    category,
	})
}

// GetCategories godoc
// @Summary Get all categories
// @Description ดึงข้อมูลหมวดหมู่ทั้งหมด
// @Tags categories
// @Produce json
// @Success 200 {array} models.JsonResponse
// @Router /api/categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to retrieve categories.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Categories retrieved successfully.",
		Data:    categories,
	})
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description แก้ไขหมวดหมู่ที่มีอยู่
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.CategoryUpdateRequest true "Category JSON"
// @Success 200 {object} models.JsonResponse
// @Router /api/categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var request models.CategoryUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse {
			Success: false,
			Message: "Invalid request data.",})
		return
	}

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Category not found.",})
		return
	}

	isExists := database.DB.Where("name = ? AND id != ?", request.Name, id).First(&models.Category{}).RowsAffected > 0
	if isExists {
		c.JSON(http.StatusConflict, models.JsonResponse{
			Success: false,
			Message: "Category with this name already exists.",
			Data:    nil,
		})
		return
	}

	category.Name = request.Name
	category.Icon = request.Icon

	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to update category.",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Category updated successfully.",
		Data:    category,
	})
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description ลบหมวดหมู่
// @Tags categories
// @Param id path int true "Category ID"
// @Success 204 {object} models.JsonResponse
// @Router /api/categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Category not found.",})
		return
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to delete category.",
			Data:    nil,
		})
		return
	}
	
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Category deleted successfully.",
		Data:    nil,
	})
}

// GetCategory godoc
// @Summary Get a single category
// @Description ดึงข้อมูลหมวดหมู่ตาม ID
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/categories/{id} [get]
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Category not found.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Category retrieved successfully.",
		Data:    category,
	})
}