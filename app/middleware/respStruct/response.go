package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("统一接口返回中间件结束")
	}
}
