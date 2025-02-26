package router

import (
	userController "ginapi/app/controller/user"
	initMiddleware "ginapi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	// 注册全局中间件
	r = initMiddleware.InitGlobleMiddleware(r)
	user := r.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}

	return r

}
