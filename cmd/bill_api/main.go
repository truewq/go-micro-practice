package main

import (
	"github.com/gin-gonic/gin"
	"go-micro-example-auth/internal/app/bill_api"
	"go-micro-example-auth/internal/pkg"
	internal "go-micro-example-auth/internal/proto"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
)

func main(){
	service := web.NewService(
		web.Name("tech.share.api.bill"),
		)

	service.Init()

	bill_api.Cli2bill = internal.NewBillService("tech.share.srv.bill", client.DefaultClient)

	router := gin.Default()
	router.Use(pkg.TokenMiddleWare)

	billHandler := new(bill_api.Handler)
	router.POST("/bill/getBills", billHandler.GetBills)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}