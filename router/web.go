package router

import (
	"fmt"
	initMiddleware "ginapi/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	// 注册全局中间件
	r = initMiddleware.InitGlobleMiddleware(r)

    // {}为了代码规范
    {
        r.GET("/ce", func(c *gin.Context) {
            // 取值
            req, _ := c.Get("request")
            fmt.Println("request111:", req)
            // 页面接收
            c.JSON(200, gin.H{"request": req})
        })

    }
	
	// userGroup := r.Group("/user")
	// userGroup.Any("/current", func(c *gin.Context) {
	// 	// user, _ := user.Current()
	// 	// c.String(http.StatusOK, "hello %s", user.Username)
	// 	user.SayHello(c)
		
	// })

	return r
	
}

