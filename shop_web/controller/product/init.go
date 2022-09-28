package product

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop/shop_product/proto/product"
)

var (
	ProductService product.ProductService
)

func init() {
	consulConf := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shop.product"),
		micro.Registry(consulConf),
	)
	ProductService = product.NewProductService("shop.product", service.Client())

}
