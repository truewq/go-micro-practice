package auth_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	internal "go-micro-example-auth/internal/proto"
	"net/http"
)

type Handler struct{}
var(
	Cli2auth internal.AuthService
)
func(*Handler) Login(c *gin.Context){
	var req internal.LoginRequest
	var res internal.LoginResponse
	var err error
	status := http.StatusOK
	defer func() {
		if err != nil{
			res.Token = ""
		}
		c.JSON(status, res)
	}()

	err = c.ShouldBindJSON(&req)
	if err != nil{
		log.Errorf("get json fail,err=%v", err)
		status = http.StatusBadRequest
		return
	}

	loginRes, err := Cli2auth.Login(context.TODO(), &req)
	if err != nil || loginRes == nil{
		res.Code = -1
		status = http.StatusInternalServerError
		return
	}
	log.Infof("loginRes=", loginRes)
	res.Token = loginRes.GetToken()
}