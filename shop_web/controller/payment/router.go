package payment

import (
	"github.com/gin-gonic/gin"
	"shop_web/middleware"
)

func Router(route *gin.RouterGroup) {
	route.GET("/list", middleware.UserJwt, PaymentList)
	route.POST("/add", middleware.UserJwt, PaymentAdd)
	route.DELETE("/del", middleware.UserJwt, PaymentDel)
}
