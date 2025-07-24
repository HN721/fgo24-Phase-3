package controller

import (
	"nastha-test/models"
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
