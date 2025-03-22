package microClient

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientConfig struct {
	Target      string
	ServiceName string
	ServicePort string
	MaxConn     int
}

var poolMap = map[string]chan *grpc.ClientConn{}

var ConfigList = []clientConfig{
	{
		Target:      "golangService:50051",
		ServiceName: "user",
		ServicePort: ":50051",
		MaxConn:     10,
	},
	{
		Target:      "golangService:50052",
		ServiceName: "order",
		ServicePort: ":50052",
		MaxConn:     10,
	},
}

func NewClientPool() {
	// 初始化客户端配置
	for _, config := range ConfigList {
		// 由于err在后续代码中会被使用，这里声明err是必要的，无需删除
		var err error
		poolMap[config.ServiceName], err = createConn(&config)
		if err != nil {
			fmt.Printf("create connection pool failed: %v", err)
			continue
		}
	}

	// 无需返回任何值，打印信息即可
	fmt.Println("client pool initialized successfully")
	fmt.Println("poolMap:", poolMap)
}

func createConn(config *clientConfig) (chan *grpc.ClientConn, error) {

	pool := make(chan *grpc.ClientConn, config.MaxConn)

	// 初始化一些空闲连接
	for i := 0; i < config.MaxConn; i++ {
		conn, err := grpc.NewClient(config.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("create connection failed: %v", err)
		}
		pool <- conn
	}
	return pool, nil
}

func GetClientConn(serverName string) (*grpc.ClientConn, error) {
	pool, ok := poolMap[serverName]
	if !ok {
		return nil, fmt.Errorf("no connection pool found for server: %s", serverName)
	}
	conn := <-pool
	return conn, nil
}

func ReleaseClientConn(serverName string, conn *grpc.ClientConn) error {
	pool, ok := poolMap[serverName]
	if !ok {
		return fmt.Errorf("no connection pool found for server: %s", serverName)
	}
	pool <- conn // 将连接放回channel中
	return nil
}
