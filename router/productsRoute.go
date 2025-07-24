package router

import (
	"nastha-test/controller"
	"nastha-test/middleware"

	"github.com/gin-gonic/gin"
)

func productsRoute(r *gin.RouterGroup) {
	r.GET("/prodcuts-cateogry", middleware.AuthMiddleware(), controller.GetAllProductCategories)
	r.GET("", middleware.AuthMiddleware(), controller.GetAllProduct)
	r.GET("/:id", middleware.AuthMiddleware(), controller.GetProductByID)
	r.POST("", middleware.AdminMiddleware(), controller.CreateProduct)
	r.PUT("/:id", middleware.AdminMiddleware(), controller.UpdateProduct)

}
