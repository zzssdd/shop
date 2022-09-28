package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop_collect/controller"
	"shop_collect/proto/collection"
)

func main() {
	consulConf := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Name("shop.collection"),
		micro.Version("latest"),
		micro.Registry(consulConf),
	)

	// Register handler
	srv.Init()

	collection.RegisterCollectionHandler(srv.Server(), new(controller.Collection))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
