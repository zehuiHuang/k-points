package main

import (
	"context"
	"fmt"
	proto "go-learn/base/protocol/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

// HelloService 业务处理服务
type HelloService struct {
	proto.UnimplementedHelloServiceServer
}

// SayHello 实现具体的业务方法逻辑
func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	return &proto.HelloResp{
		Reply: fmt.Sprintf("hello name: %s", req.Name),
	}, nil
}

func main() {
	// 创建 tcp 端口监听器
	listener, err := net.Listen("tcp", ":8093")
	if err != nil {
		panic(err)
	}

	// 创建 grpc server
	server := grpc.NewServer()
	// 将自定义的业务处理服务注册到 grpc server 中
	proto.RegisterHelloServiceServer(server, &HelloService{})
	// 运行 grpc server
	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("end----------------------------")
}
