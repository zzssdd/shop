package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/shop_user/proto/user"
	"shop_web/utils"
)

func Register(c *gin.Context) {
	email := c.PostForm("email")
	is_ok := utils.VerifyEmail(email)
	if !is_ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
		return
	}
	password := c.PostForm("password")
	captche := c.PostForm("code")
	register_data := &user.RegistryRequest{
		Email:    email,
		Password: password,
		Code:     captche,
	}
	rep, _ := UserService.Registry(context.TODO(), register_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}

func UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	is_ok := utils.VerifyEmail(email)
	if !is_ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
		return
	}
	login_data := &user.LoginRequest{
		Email:    email,
		Password: password,
	}
	rep, _ := UserService.Login(context.TODO(), login_data)
	userToken, err := utils.NewToken(email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "生成token错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  rep.Code,
		"msg":   rep.Msg,
		"token": userToken,
		"user":  email,
	})
}

func SendEmail(c *gin.Context) {
	email := c.PostForm("email")
	is_ok := utils.VerifyEmail(email)
	fmt.Println(email)
	if !is_ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
		return
	}
	email_data := &user.EmailRequest{
		Email: email,
	}
	rep, _ := UserService.SendEmail(context.TODO(), email_data)
	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})
}
