package router

import (
	"nastha-test/controller"

	"github.com/gin-gonic/gin"
)

func productsRoute(r *gin.RouterGroup) {
	r.GET("/prodcuts-cateogry", controller.GetAllProductCategories)
	r.GET("", controller.GetAllProduct)

}
