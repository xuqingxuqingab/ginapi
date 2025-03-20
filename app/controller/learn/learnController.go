package learn

import (
	"ginapi/app/gen/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Learn1(c *gin.Context) {
	// 错误信息表明在结构体字面量中使用了无效的字段名 "Id"，推测结构体字段名可能为 "ID" 或其他。
	// 假设正确的字段名为 "ID"，修改后的代码如下：
	res := user.UserInfoRequest{}
	res.Id = 1
	// 修复：使用指针传递以避免复制包含锁的值
	c.JSON(http.StatusOK, &res)
}
