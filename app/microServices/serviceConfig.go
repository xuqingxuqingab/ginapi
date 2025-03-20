package microServices

import (
	"fmt"
	"ginapi/app/gen/user"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

// 微服务配置
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

type server struct {
	user.UnimplementedUserServiceServer
}

func InitializeMicroServices() {
	var wg sync.WaitGroup

	for _, serviceConfig := range ConfigList {
		wg.Add(1)
		go func(config MicroServicesConfig) {
			defer wg.Done()
			fmt.Println("启动微服务：", config.Name)
			InitializeGrpcServer(config.Port)
		}(serviceConfig)
	}
	fmt.Println("所有微服务已启动")
	wg.Wait()
}
func InitializeGrpcServer(port string) {
	// lis, err := net.Listen("tcp", ":50051")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	log.Println("Server started on " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
