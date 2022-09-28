package router

import (
	"github.com/gin-gonic/gin"
	"shop_web/controller/collect"
	"shop_web/controller/payment"
	"shop_web/controller/product"
	"shop_web/controller/seckill"
	"shop_web/controller/user"
	"shop_web/middleware"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.Cors())
	user_Group := router.Group("/user")
	product_Group := router.Group("/product")
	seckill_Group := router.Group("/seckill")
	collect_Group := router.Group("/collect")
	payment_Group := router.Group("/payment")

	user.Router(user_Group)
	product.Router(product_Group)
	seckill.Router(seckill_Group)
	collect.Router(collect_Group)
	payment.Router(payment_Group)

	router.Static("/upload", "./upload")
}
