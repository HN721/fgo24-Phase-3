package controller

import (
	"nastha-test/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name          string  `json:"name"`
	Image         string  `json:"image"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	Stock         int     `json:"stock"`
	CategoryID    int     `json:"category_id"`
}

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
func GetProductByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid product ID", "error": err.Error()})
		return
	}

	data, err := models.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Product not found", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully fetched product by ID",
		"result":  data,
	})
}

func CreateProduct(ctx *gin.Context) {
	var req ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
		return
	}

	product := models.Products{
		Name:           req.Name,
		Image:          req.Image,
		Purchase_price: req.PurchasePrice,
		Selling_price:  req.SellingPrice,
		Stock:          req.Stock,
	}

	err := models.CreateProduct(product, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create product", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "message": "Product created successfully"})
}
func UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid product ID", "error": err.Error()})
		return
	}

	var req ProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid request body", "error": err.Error()})
		return
	}

	product := models.Products{
		Name:           req.Name,
		Image:          req.Image,
		Purchase_price: req.PurchasePrice,
		Selling_price:  req.SellingPrice,
		Stock:          req.Stock,
	}

	err = models.UpdateProduct(id, product, req.CategoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to update product", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Product updated successfully"})
}
