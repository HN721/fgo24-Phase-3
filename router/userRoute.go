package router

import (
	"nastha-test/controller"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.POST("/register", controller.CreateUsers)
}
