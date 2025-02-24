package router

import (
	userController "ginapi/app/controller/user"
	initMiddleware "ginapi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	// 注册全局中间件
	r = initMiddleware.InitGlobleMiddleware(r)

	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/SayHello", userController.SayHello)

		v1 = v1.Group("/user")
		{
			v1.GET("/login", userController.Login)
		}
	}

	v2 := r.Group("/v2")
	// {} 是书写规范
	{
		v2.GET("/SayHello", userController.SayHello2)

		v2 = v2.Group("/user")
		{
			v2.GET("/login", userController.Login2)
		}
	}

	return r

}
