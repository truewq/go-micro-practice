package pkg

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2/util/log"
)

const(
	tokenSecretKey = "keyOnlyServerShouldKnow"
)

var(
	idGen, _ = snowflake.NewNode(1)
)

func CreateToken(user string, password string)(string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id": idGen.Generate().String(),
		"NotBefore": time.Now().Unix(),
		"ExpiresAt": time.Now().Add(time.Hour*2).Unix(),
		"Issuer": "tech share",
		"User": user,
	})

	return token.SignedString([]byte(tokenSecretKey))
}

func TokenMiddleWare(c *gin.Context){
	cliToken := c.GetHeader("Authorization")
	log.Infof("cliToken=", cliToken)
	tokenData, err := CheckToken(cliToken)
	if err != nil || tokenData == nil{
		c.JSON(http.StatusBadRequest, "check token fail")
		c.Abort()
		return
	}

	log.Infof("user=", tokenData["User"])
	c.Set("User", tokenData["User"])
	c.Next()
}


func CheckToken(clientToken string) (jwt.MapClaims, error){
	token, err := jwt.Parse(clientToken, func(token *jwt.Token)(interface{}, error){
		return []byte(tokenSecretKey), nil
	})

	if err != nil{
		return nil, fmt.Errorf("pass token fail=%v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		return nil, fmt.Errorf("trans claim fail")
	}

	return claims, nil
}