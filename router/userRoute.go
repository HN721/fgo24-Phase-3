package router

import (
	"nastha-test/controller"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	r.POST("/register", controller.CreateUsers)
	r.POST("/login", controller.LoginUser)

}
