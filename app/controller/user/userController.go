package user

import (
	"net/http"

	"ginapi/app/api/user"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) *user.SayHelloRes {
	
	res := user.SayHelloRes{
		Message: "hello hello",
	}

	c.JSON(http.StatusOK, res)
	
	return &res
}