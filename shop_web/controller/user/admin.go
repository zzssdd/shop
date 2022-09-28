package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_user/proto/admin"
	"shop_web/utils"
	"strconv"
)

func UserList(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.DefaultQuery("currentPage", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	list_data := &admin.ListRequest{
		CurrentPage: int32(currentPage),
		PageSize:    int32(pageSize),
	}
	rep, err := AdminService.UserList(context.TODO(), list_data)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"users": rep.User,
		"total": rep.Total,
	})
}

func AdminLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	login_data := &admin.LoginRequest{
		Name:     username,
		Password: password,
	}
	rep, err := AdminService.Login(context.TODO(), login_data)
	adminToken, err := utils.NewToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"token": adminToken,
		"user":  username,
	})
}
