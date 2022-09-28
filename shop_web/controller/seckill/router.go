package seckill

import (
	"github.com/gin-gonic/gin"
	"shop_web/middleware"
)

func Router(route *gin.RouterGroup) {
	route.GET("/list", SeckillList)
	route.POST("/add", middleware.AdminJwt, SeckillAdd)
	route.DELETE("/del", middleware.AdminJwt, SeckillDel)
	route.GET("/info", SeckillInfo)
	route.PUT("/edit", middleware.AdminJwt, SeckillEdit)

	route.POST("/seckill", middleware.UserJwt, Seckill)

	route.GET("result", middleware.UserJwt, SeckillResult)
}
