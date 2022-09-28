package controller

import (
	"context"
	"github.com/patrickmn/go-cache"
	"shop_user/model"
	"shop_user/proto/user"
	"shop_user/utils"
	"time"
)

type User struct {
}

var c = cache.New(60*time.Second, 60*time.Second)

func (u *User) Registry(ctx context.Context, in *user.RegistryRequest, out *user.Response) error {
	email := in.Email
	password := utils.PWMd5(in.Password)
	capthch := in.Code

	code, is_ok := c.Get(email)
	if !is_ok {
		out.Code = 500
		out.Msg = "未获取验证码"
		return nil
	}
	if code != capthch {
		out.Code = 500
		out.Msg = "验证码输入不正确,请重新输入"
		return nil
	}
	err := model.UserRegistry(email, password)
	if err != nil {
		out.Code = 500
		out.Msg = "用户注册失败"
		return nil
	}
	out.Code = 200
	out.Msg = "用户注册成功"
	return nil
}
func (u *User) Login(ctx context.Context, in *user.LoginRequest, out *user.Response) error {
	email := in.Email
	password := utils.PWMd5(in.Password)
	is_ok := model.UserLogin(email, password)
	if !is_ok {
		out.Code = 500
		out.Msg = "账号密码输入错误"
		return nil
	}
	out.Code = 200
	out.Msg = "登录成功"
	return nil
}
func (u *User) SendEmail(ctx context.Context, in *user.EmailRequest, out *user.Response) error {
	email := in.Email
	is_ok := model.UserExistEmail(email)
	if !is_ok {
		out.Code = 500
		out.Msg = "该邮箱已被注册"
		return nil
	}
	code := utils.RandNum(6)
	c.Set(email, code, cache.DefaultExpiration)
	err := utils.SendEmail(email, code)
	if err != nil {
		out.Code = 500
		out.Msg = "发送验证码失败，请稍后再试"
		return err
	}
	out.Code = 200
	out.Msg = "发送验证码成功"
	return nil
}
