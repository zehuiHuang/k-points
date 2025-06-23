package main

import (
	"context"
	"fmt"
	proto "go-learn/base/protocol/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 通过指定地址，建立与 grpc 服务端的连接
	conn, err := grpc.Dial("localhost:8093", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:" + err.Error())
		return
	}
	// ...
	// 调用 .grpc.pb.go 文件中预生成好的客户端构造器方法，创建 grpc 客户端
	client := proto.NewHelloServiceClient(conn)

	// 调用 .grpc.pb.go 文件预生成好的客户端请求方法，使用 .pb.go 文件中预生成好的请求参数作为入参，向 grpc 服务端发起请求
	resp, err := client.SayHello(context.Background(), &proto.HelloReq{
		Name: "xiaoxuxiansheng",
	})
	// ...
	// 打印取得的响应参数
	fmt.Printf("resp: %+v", resp)
}
