package user

import (
	"github.com/gin-gonic/gin"
	"shop_web/middleware"
)

func Router(r *gin.RouterGroup) {
	r.POST("/register", Register)
	r.POST("/email", SendEmail)
	r.POST("/login", UserLogin)

	r.POST("/admin_login", AdminLogin)
	r.GET("/list", middleware.AdminJwt, UserList)
}
