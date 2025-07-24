package controller

import (
	"nastha-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProductCategories(ctx *gin.Context) {
	data, err := models.GetAllProductCategory()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"succes":  false,
			"message": "Cannot Get Data From Database",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"succes":  true,
		"message": "Successsfully Get Products Category",
		"results": data,
	})
}

func GetAllProduct(ctx *gin.Context) {
	data, err := models.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"succes":  false,
			"message": "Cannot Get Data From Database",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"succes":  true,
		"message": "Successsfully Get Products",
		"results": data,
	})
}
