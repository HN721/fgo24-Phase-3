package router

import (
	"nastha-test/controller"

	"github.com/gin-gonic/gin"
)

func transactionRoute(r *gin.RouterGroup) {
	r.POST("", controller.CreateTransactionCtrl)
	r.GET("", controller.GetTransactionHistoryCtrl)

}
