package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop_user/controller"
	"shop_user/proto/admin"
	"shop_user/proto/user"
)

func main() {
	// Create service
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	srv := micro.NewService(
		micro.Name("shop.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
	)
	srv.Init()
	// Register handler
	admin.RegisterAdminHandler(srv.Server(), new(controller.Admin))
	user.RegisterUserHandler(srv.Server(), new(controller.User))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
