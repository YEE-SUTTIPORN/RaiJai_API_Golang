package controllers

import (
	"RaiJaiAPI_Golang/database"
	"RaiJaiAPI_Golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description เพิ่มรายการธุรกรรม
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body models.TransactionCreateRequest true "Transaction JSON"
// @Success 201 {object} models.JsonResponse
// @Router /api/transactions [post]
func CreateTransaction(c *gin.Context) {
	var request models.TransactionCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{
			Success: false,
			Message: "Invalid request data.",
			Data:    nil,
		})
		return
	}

	transaction := models.Transaction{
		Amount:     request.Amount,
		Note:       request.Note,
		Date:       request.Date,
		UserID:     request.UserID,
		CategoryID: request.CategoryID,
	}

	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to create transaction.",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, models.JsonResponse{
		Success: true,
		Message: "Transaction created successfully.",
		Data:    transaction,
	})
}

// GetTransactions godoc
// @Summary Get all transactions
// @Description ดึงข้อมูลธุรกรรมทั้งหมด
// @Tags transactions
// @Produce json
// @Success 200 {array} models.JsonResponse
// @Router /api/transactions [get]
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	if err := database.DB.Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to retrieve transactions.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Transactions retrieved successfully.",
		Data:    transactions,
	})
}

// UpdateTransaction godoc
// @Summary Update an existing transaction
// @Description แก้ไขธุรกรรมที่มีอยู่
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param transaction body models.TransactionUpdateRequest true "Transaction JSON"
// @Success 200 {object} models.JsonResponse
// @Router /api/transactions/{id} [put]
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var request models.TransactionUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.JsonResponse{
			Success: false,
			Message: "Invalid request data.",
			Data:    nil,
		})
		return
	}

	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Transaction not found.",
			Data:    nil,
		})
		return
	}

	transaction.Amount = request.Amount
	transaction.Note = request.Note
	transaction.Date = request.Date
	transaction.UserID = request.UserID
	transaction.CategoryID = request.CategoryID
	if err := database.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to update transaction.",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Transaction updated successfully.",
		Data:    transaction,
	})
}

// DeleteTransaction godoc
// @Summary Delete a transaction
// @Description ลบธุรกรรม
// @Tags transactions
// @Param id path int true "Transaction ID"
// @Success 204 {object} models.JsonResponse
// @Router /api/transactions/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Transaction not found.",
			Data:    nil,
		})
		return
	}
	if err := database.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.JsonResponse{
			Success: false,
			Message: "Failed to delete transaction.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusNoContent, models.JsonResponse{
		Success: true,
		Message: "Transaction deleted successfully.",
		Data:    nil,
	})
}

// GetTransaction godoc
// @Summary Get a transaction by ID
// @Description ดึงข้อมูลธุรกรรมตาม ID
// @Tags transactions
// @Param id path int true "Transaction ID"
// @Success 200 {object} models.JsonResponse
// @Router /api/transactions/{id} [get]
func GetTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.JsonResponse{
			Success: false,
			Message: "Transaction not found.",
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, models.JsonResponse{
		Success: true,
		Message: "Transaction retrieved successfully.",
		Data:    transaction,
	})
}