package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
	Result bool   `json:"result"`
}

// 定义中间
func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Next()
        fmt.Println("统一接口返回中间件结束")
    }
}

func ApiReturn(data any, err error) ApiResponse {


	fmt.Println(data)
	fmt.Println(err)

	// // 如果获取到error 就返回失败
	if err != nil {
		return Fail(err.Error())
	} else {
		// 否则返回成功
		return Success(data)
	}
}

func Success(data any) ApiResponse {
	return ApiResponse{
		Code: 200,
		Msg:  "success",
		Data: data,
		Result: true,
	}
}

func Fail(msg string) ApiResponse {
	return ApiResponse{
		Code: 400,
		Msg:  msg,
		Data: nil,
		Result: false,
	}
}