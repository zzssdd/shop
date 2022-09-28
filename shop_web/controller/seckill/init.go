package seckill

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop/shop_product/proto/seckillManager"
	"shop/shop_seckill/proto"
)

var (
	ManagerService seckillManager.SeckillsService
	SeckillService proto.DoSeckillService
)

func init() {
	consulConf := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shop.seckill"),
		micro.Registry(consulConf),
	)
	ManagerService = seckillManager.NewSeckillsService("shop.product", service.Client())
	SeckillService = proto.NewDoSeckillService("shop.seckill", service.Client())
}
