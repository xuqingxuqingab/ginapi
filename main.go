package main

import (
	"ginapi/app/bootstrap"
	"ginapi/app/global"
	"ginapi/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化config文件
	bootstrap.InitializeConfig()

	// 初始化数据库
	dbs := bootstrap.InitializeDB()
	global.App.Dbs = dbs

	// 创建路由
	r := gin.Default()
	// 绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	router.InitRouter(r)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	port := global.App.Config.App.Port
	r.Run(":" + port)
}
