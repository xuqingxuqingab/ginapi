package microServices

import (
	"context"
	"ginapi/app/gen/user"
)

// 定义一个用户详情服务数据接口
func (s *userServer) GetUserInfo(ctx context.Context, req *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 修改结构体字面量的字段赋值语法，正确的语法应该使用 . 来指定字段名
	return &user.UserInfoResponse{Id: 1, Name: "xuqing"}, nil
}

// 定义一个用户列表服务数据接口
func (s *userServer) GetUserList(ctx context.Context, req *user.UserListRequest) (*user.UserListResponse, error) {
	// 假设你有一个名为 users 的切片，其中包含多个 UserInfoResponse 结构体
	users := []*user.User{
		{Id: 1, Name: "xuqing"},
		{Id: 2, Name: "xuqing2"},
		{Id: 3, Name: "xuqing3"},
	}
	// 创建一个新的 UserListResponse 结构体，并将 users 切片赋值给它的 Users 字段
	response := &user.UserListResponse{
		Users: users,
		Total: 1,
	}
	return response, nil
}
