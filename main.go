package main

import (
	"fmt"
	"ginapi/app/bootstrap"
	"ginapi/app/global"
	microServices "ginapi/app/microServices"
	"ginapi/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	projectType := os.Args[1:]

	// 初始化config文件
	bootstrap.InitializeConfig()

	// 初始化数据库
	dbs := bootstrap.InitializeDB()
	global.App.Dbs = dbs

	fmt.Println(projectType)
	if len(projectType) == 0 { // 认为是客户端启动
		fmt.Println("启动客户端")

		// 创建路由
		r := gin.Default()
		// 绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		router.InitRouter(r)

		// 3.监听端口，默认在8080
		// Run("里面不指定端口号默认为8080")
		port := global.App.Config.App.Port
		r.Run(":" + port)

	} else { // 启动grpc服务端
		fmt.Println("启 动 grpc服务端")
		microServices.InitializeMicroServices()
	}
}
