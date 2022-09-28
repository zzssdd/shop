package controller

import (
	"context"
	"shop_user/model"
	"shop_user/proto/admin"
	"shop_user/utils"
)

type Admin struct {
}

func (a *Admin) Login(ctx context.Context, request *admin.LoginRequest, response *admin.Response) error {
	name := request.Name
	password := utils.PWMd5(request.Password)
	is_ok := model.AdminCheckLogin(name, password)
	if !is_ok {
		response.Code = 500
		response.Msg = "账号或密码错误"
		return nil
	}
	response.Code = 200
	response.Msg = "登录成功"
	return nil
}
func (a *Admin) UserList(ctx context.Context, request *admin.ListRequest, response *admin.ListResponse) error {
	currentPage := request.CurrentPage
	pageSize := request.PageSize
	users, count, err := model.UserList(int(currentPage), int(pageSize))
	if err != nil {
		response.Code = 500
		response.Msg = "获取用户列表失败"
		return nil
	}
	rep_users := []*admin.User{}
	for _, user := range users {
		rep_user := &admin.User{
			Id:         int32(user.ID),
			Email:      user.Email,
			Status:     int32(user.Status),
			CreateTime: user.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		rep_users = append(rep_users, rep_user)
	}
	response.Code = 200
	response.Msg = "获取列表成功"
	response.User = rep_users
	response.Total = int32(count)
	return nil
}
func (a *Admin) UserDel(ctx context.Context, request *admin.DelRequest, response *admin.Response) error {
	id := request.Id
	err := model.UserDel(int(id))
	if err != nil {
		response.Code = 500
		response.Msg = "删除失败，请重试"
		return nil
	}
	response.Code = 200
	response.Msg = "删除成功"
	return nil
}
