package user

import (
	"net/http"

	"ginapi/app/api/user"

	userBussiness "ginapi/app/business/user"
	"ginapi/app/common"
	"ginapi/app/model/test1"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := user.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查账号是否存在
	isExist := userBussiness.CheckAccountIsExist(req.Account)
	if !isExist {
		common.JsonToFront(c, &common.ApiResponse{
			Code:   400,
			Msg:    "账号不存在",
			Data:   req,
			Result: false,
		})
		return
	}
	// 验证账号密码
	isCorrect := userBussiness.CheckPassword(req.Account, req.Password)

	if !isCorrect {
		common.JsonToFront(c, &common.ApiResponse{
			Code:   400,
			Msg:    "账号或密码错误",
			Data:   req,
			Result: false,
		})
		return
	}

	common.JsonToFront(c, &common.ApiResponse{
		Code:   200,
		Msg:    "欢迎" + req.Account,
		Data:   req,
		Result: true,
	})

}

func Register(c *gin.Context) {

	// 获取请求参数
	req := user.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 检测账号是否已注册
	isExist := userBussiness.CheckAccountIsExist(req.Account)
	if isExist {
		common.JsonToFront(c, &common.ApiResponse{
			Code:   400,
			Msg:    "账号已存在",
			Data:   req,
			Result: false,
		})
		return
	}
	// 生成密码盐
	salt := "abcde"

	// 创建用户信息
	userBussiness.CreateUser(&test1.TestUser{
		Name:     req.Name,
		Password: req.Password,
		Salt:     salt,
		Account:  req.Account,
	})
	common.JsonToFront(c, &common.ApiResponse{
		Code:   200,
		Msg:    "注册成功",
		Data:   req,
		Result: true,
	})
}
