package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"go-micro-example-auth/internal/app/auth_srv"
	internal "go-micro-example-auth/internal/proto"
)

func main(){
	service := micro.NewService(
		micro.Name("tech.share.srv.auth"),
		)

	service.Init()

	err := internal.RegisterAuthHandler(service.Server(), new(auth_srv.Handler))
	if err != nil{
		log.Fatal("register auth handler fail")
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
