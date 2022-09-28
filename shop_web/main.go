package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"shop_web/router"
)

func main() {
	route := gin.Default()

	router.InitRouter(route)

	service := web.NewService(
		web.Name("shop.web"),
		web.Version("latest"),
		web.Handler(route),
		web.Address(":8081"),
	)
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
