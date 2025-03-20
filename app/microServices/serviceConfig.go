package microServices

import (
	"fmt"
	"ginapi/app/gen/order" // 确保 order 服务实现存在
	"ginapi/app/gen/user"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type MicroServicesConfig struct {
	Name string
	Port string
}

var ConfigList = []MicroServicesConfig{
	{
		Name: "user",
		Port: ":50051",
	},
	{
		Name: "order",
		Port: ":50052",
	},
}

type userServer struct {
	user.UnimplementedUserServiceServer
}

type orderServer struct {
	order.UnimplementedUserServiceServer
}

func InitializeMicroServices() {
	// 启动所有微服务
	for _, config := range ConfigList {
		go func(cfg MicroServicesConfig) {
			log.Printf("正在启动微服务: %s", cfg.Name)
			if err := InitializeGrpcServer(cfg); err != nil {
				log.Printf("微服务 %s 启动失败: %v", cfg.Name, err)
			}
		}(config)
	}

	// 等待终止信号实现优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("收到终止信号，正在关闭服务...")
}

func InitializeGrpcServer(config MicroServicesConfig) error {
	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		return fmt.Errorf("端口监听失败: %v", err)
	}

	server := grpc.NewServer()

	// 根据服务名注册不同实现
	switch config.Name {
	case "user":
		user.RegisterUserServiceServer(server, &userServer{})
	case "order":
		order.RegisterUserServiceServer(server, &orderServer{})
	default:
		return fmt.Errorf("未知服务类型: %s", config.Name)
	}

	log.Printf("微服务 %s 已启动，监听端口 %s", config.Name, config.Port)
	if err := server.Serve(lis); err != nil {
		return fmt.Errorf("服务运行失败: %v", err)
	}
	return nil
}
