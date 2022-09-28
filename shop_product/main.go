package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop_product/controller"
	"shop_product/proto/product"
	"shop_product/proto/seckillManager"
)

func main() {
	// Create service
	consulConf := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	srv := micro.NewService(
		micro.Name("shop.product"),
		micro.Version("latest"),
		micro.Registry(consulConf),
	)
	srv.Init()
	// Register handler
	product.RegisterProductHandler(srv.Server(), new(controller.Product))
	seckillManager.RegisterSeckillsHandler(srv.Server(), new(controller.Seckill))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
