package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop_seckill/controller"
	"shop_seckill/proto"
)

func main() {
	consulConf := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Name("shop_seckill"),
		micro.Registry(consulConf),
	)

	// Register handler
	srv.Init()
	proto.RegisterDoSeckillHandler(srv.Server(), new(controller.Seckill))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
