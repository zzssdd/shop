package collect

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop/shop_collect/proto/collection"
)

var (
	CollectServer collection.CollectionService
)

func init() {
	consulConf := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shop.collection"),
		micro.Registry(consulConf),
	)
	CollectServer = collection.NewCollectionService("shop.collection", service.Client())
}
