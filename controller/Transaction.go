package controller

import (
	"nastha-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionCtrl(ctx *gin.Context) {
	var trx models.Transactions

	if err := ctx.ShouldBindJSON(&trx); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if err := models.CreateTransaction(trx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create transaction",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Transaction created successfully",
	})
}
func GetTransactionHistoryCtrl(ctx *gin.Context) {
	history, err := models.GetTransactionHistory()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get transaction history",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Transaction history fetched successfully",
		"data":    history,
	})
}
