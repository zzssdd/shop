package user

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop/shop_user/proto/admin"
	"shop/shop_user/proto/user"
)

var (
	UserService  user.UserService
	AdminService admin.AdminService
)

func init() {
	registry := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shop.user"),
		micro.Registry(registry),
	)
	UserService = user.NewUserService("shop.user", service.Client())
	AdminService = admin.NewAdminService("shop.user", service.Client())
}
