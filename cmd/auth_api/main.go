package main

import (
	"github.com/gin-gonic/gin"
	"go-micro-example-auth/internal/app/auth_api"
	internal "go-micro-example-auth/internal/proto"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func main(){
	service := web.NewService(
		web.Name("tech.share.api.auth"),
		web.Registry(kubernetes.NewRegistry()),
		web.Address(":9200"),
		)

	service.Init()

	auth_api.Cli2auth = internal.NewAuthService("tech.share.srv.auth", client.DefaultClient)

	authHandler := new(auth_api.Handler)
	router := gin.Default()
	router.POST("/auth/login", authHandler.Login)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}