package router

import (
	"nastha-test/controller"
	"nastha-test/middleware"

	"github.com/gin-gonic/gin"
)

func transactionRoute(r *gin.RouterGroup) {
	r.POST("", middleware.AdminMiddleware(), controller.CreateTransactionCtrl)
	r.GET("", middleware.AdminMiddleware(), controller.GetTransactionHistoryCtrl)

}
