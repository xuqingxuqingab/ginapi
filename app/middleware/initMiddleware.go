package initMiddleware

import (
	response "ginapi/app/middleware/respStruct"
	"ginapi/app/middleware/test"
	"ginapi/app/middleware/test1"

	"github.com/gin-gonic/gin"
)

// 注册全局中间件
func registerGlobalMiddleware() []gin.HandlerFunc {
	// 声明一个全局中间件数组配置，在初始化的时候循环调用
	middleware := []gin.HandlerFunc{
		// 统一返回中间件
		response.MiddleWare(),
		// 跨域中间件
		test.MiddleWare(),
		test1.MiddleWare(),
	}
	return middleware
}

func InitGlobleMiddleware(r *gin.Engine) (*gin.Engine) {
	middlewares := registerGlobalMiddleware()

	// 循环注册全局中间件
	for _, m := range middlewares {
		r.Use(m)
	}
	return r
}


