package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"go-micro-example-auth/internal/app/bill_srv"
	internal "go-micro-example-auth/internal/proto"
)

func main(){
	service := micro.NewService(
		micro.Name("tech.share.srv.bill"),
		)

	service.Init()

	err := internal.RegisterBillHandler(service.Server(), new(bill_srv.Handler))
	if err != nil{
		log.Fatal("register bill handler fail")
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
