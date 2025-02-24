package user

import (
	"fmt"
	"net/http"

	"ginapi/app/api/user"
	"ginapi/app/global"

	"ginapi/app/model/ppclibrary"

	"github.com/gin-gonic/gin"
)

func SayHello(c *gin.Context) {

	// 查询Resource
	db := global.App.Dbs["ppc_library"]

	// result := db.First(&ppclibrary.Resource{}, "resource_id = ?", 1)
	resource := &ppclibrary.Resource{}
	db.First(resource)
	// 打印整个结构体
	fmt.Println("-----start-------")
	fmt.Printf("Resource: %+v\n", resource)
	fmt.Println(resource.CreatedAt)
	fmt.Println(resource.UpdatedAt)
	fmt.Println(resource.Name)

	fmt.Println("-----end-------")
}

func SayHello2(c *gin.Context) {

	res := user.SayHelloRes{
		Message: "hello hello 2",
	}

	c.JSON(http.StatusOK, res)
	return
}

func Login(c *gin.Context) {

	res := user.SayHelloRes{
		Message: "hello hello login 1",
	}

	c.JSON(http.StatusOK, res)
	return
}

func Login2(c *gin.Context) {

	res := user.SayHelloRes{
		Message: "hello hello login 2",
	}

	c.JSON(http.StatusOK, res)
	return
}
