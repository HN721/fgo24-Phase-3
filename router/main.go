package router

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	userRoute(r.Group("/auth"))
	productsRoute(r.Group("/product"))
	transactionRoute(r.Group("/trx"))
}
