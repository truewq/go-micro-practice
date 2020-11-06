package auth_srv

import (
	"context"
	"go-micro-example-auth/internal/pkg"
	"github.com/micro/go-micro/v2/util/log"
	internal "go-micro-example-auth/internal/proto"
)
type Handler struct {
}

func (*Handler) Login(ctx context.Context,
	in *internal.LoginRequest,
	out *internal.LoginResponse) error{
	user := in.GetUser()
	password := in.GetPassword()
	log.Infof("user=,password=", user, password)
	if user != "tech" || password != "share"{
		out.Code = -1
		out.Token = ""
		return nil
	}

	tokenString, err := pkg.CreateToken(user, password)
	if err != nil{
		log.Errorf("create token err=", err)
		out.Code = -1
		out.Token = ""
		return nil
	}

	out.Code = 0
	out.Token = tokenString
	return nil
}

