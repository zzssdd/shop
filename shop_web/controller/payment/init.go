package payment

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"shop/shop_payment/proto/payment"
)

var (
	PaymentServer payment.PaymentService
)

func init() {
	consulConf := consul.NewRegistry()
	service := micro.NewService(
		micro.Name("shop.payment"),
		micro.Registry(consulConf),
	)
	PaymentServer = payment.NewPaymentService("shop.payment", service.Client())
}
