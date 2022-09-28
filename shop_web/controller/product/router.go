package product

import (
	"github.com/gin-gonic/gin"
	"shop_web/middleware"
)

func Router(router *gin.RouterGroup) {
	router.GET("/list", ProductList)
	router.POST("/add", middleware.AdminJwt, ProductAdd)
	router.DELETE("/del", middleware.AdminJwt, ProductDel)
	router.GET("/info", ProductInfo)
	router.PUT("/edit", middleware.AdminJwt, ProductEdit)
}
