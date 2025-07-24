package controller

import (
	"nastha-test/models"
	"nastha-test/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsers(ctx *gin.Context) {
	var req models.Users

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	err = models.SaveUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "OK",
		"results": req,
	})
}
func LoginUser(ctx *gin.Context) {
	var req models.Users

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	user, err := models.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}
