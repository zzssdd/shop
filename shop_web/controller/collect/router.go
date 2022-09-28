package collect

import (
	"github.com/gin-gonic/gin"
	"shop_web/middleware"
)

func Router(route *gin.RouterGroup) {
	route.GET("/list", middleware.UserJwt, CollectList)
	route.POST("/add", middleware.UserJwt, CollectAdd)
	route.DELETE("/del", middleware.UserJwt, CollectDel)
}
