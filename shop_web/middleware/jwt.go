package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop_web/utils"
	"strings"
)

func AdminJwt(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")

	if auth_header == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	auths := strings.Split(auth_header, " ")
	bearer := auths[0]
	token := auths[1]
	if len(token) == 0 || len(bearer) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	user, err := utils.AuthToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	c.Set("admin_user", user.UserName)
	c.Next()
}
func UserJwt(c *gin.Context) {
	auth_header := c.Request.Header.Get("Authorization")
	if auth_header == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	auth := strings.Split(auth_header, " ")

	bearer := auth[0]
	token := auth[1]
	if len(token) == 0 || len(bearer) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	user, err := utils.AuthToken(token)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请先登录",
		})
		c.Abort()
		return
	}
	c.Set("user", user.UserName)
	c.Next()
}
