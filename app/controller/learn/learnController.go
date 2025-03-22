package learn

import (
	"ginapi/app/gen/user"
	"ginapi/app/microClient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Learn1(c *gin.Context) {
	// 错误信息表明在结构体字面量中使用了无效的字段名 "Id"，推测结构体字段名可能为 "ID" 或其他。
	// 假设正确的字段名为 "ID"，修改后的代码如下：
	// res := user.UserInfoRequest{}
	// res.Id = 1
	// // 修复：使用指针传递以避免复制包含锁的值
	// c.JSON(http.StatusOK, &res)

	// 获取微服务的客户端连接池
	conn, _ := microClient.GetClientConn("user")
	defer microClient.ReleaseClientConn("user", conn)

	client := user.NewUserServiceClient(conn)

	// 调用微服务的方法
	// res, err := client.GetUserInfo(c, &user.UserInfoRequest{Id: 1})
	res, err := client.GetUserList(c, &user.UserListRequest{Page: 1, PageSize: 10})

	c.JSON(http.StatusOK, res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
