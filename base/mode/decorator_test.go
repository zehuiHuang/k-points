package mode

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	rice := NewRice()
	rice.Eat()
	rice.Cost()
	fmt.Println(rice.Eat())
	fmt.Println(rice.Cost())

	rice = NewLaoGanMaDecorator(rice)
	fmt.Println(rice.Eat())
	fmt.Println(rice.Cost())
}

func TestChainUnaryInterceptors(t *testing.T) {
	var myInterceptor = func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error) {
		// 添加前处理...
		fmt.Printf("interceptor preprocess, req: %+v\n", req)
		resp, err = handler(ctx, req)
		// 添加后处理...
		fmt.Printf("interceptor postprocess, req: %+v\n", resp)
		return
	}
	var myInterceptor2 = func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error) {
		// 添加前处理...
		fmt.Printf("interceptor preprocess2, req: %+v\n", req)
		resp, err = handler(ctx, req)
		// 添加后处理...
		fmt.Printf("interceptor postprocess2, req: %+v\n", resp)
		return
	}
	var chainUnaryInts = []UnaryServerInterceptor{myInterceptor, myInterceptor2}
	f := chainUnaryInterceptors(chainUnaryInts)
	f(context.Background(), nil, &UnaryServerInfo{}, func(ctx context.Context, req any) (any, error) {
		fmt.Printf("ddddddddddddd\n")
		return func() {
			fmt.Printf("ffffffffffff")
		}, errors.New("1111111111111")
	})
}

func TestDecorate(t *testing.T) {
	f := Decorate(func(ctx context.Context, param map[string]interface{}) error {
		fmt.Println("dddddddd")
		return nil
	})
	var param map[string]interface{}
	err := f(context.Background(), param)
	if err != nil {
		return
	}
}

func TestWrapToolCall2(t *testing.T) {
	f := wrapToolCall2(&Weather2{}, nil, true)

	data := "testdata"
	result, err := f(context.Background(), &data)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Result:", *result)
}

// -----------------------------------------------------------------------test-------------VVVVVVVVVVVVVVVVVVVVVVVVVV待删除
//var myInterceptor1 = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//	fmt.Printf("interceptor1 preprocess, req: %+v\n", req)
//	resp, err = handler(ctx, req)
//	fmt.Printf("interceptor1 postprocess, req: %+v\n", resp)
//	return
//}
//
//var myInterceptor2 = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//	fmt.Printf("interceptor2 preprocess, req: %+v\n", req)
//	resp, err = handler(ctx, req)
//	fmt.Printf("interceptor2 postprocess, resp: %+v\n", resp)
//	return
//}
//
//func (s *Server) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
//	fmt.Println("core handle logic......")
//	return &proto.HelloResp{
//		Reply: fmt.Sprintf("hello name: %s", req.Name),
//	}, nil
//}
//
//func main() {
//	listener, err := net.Listen("tcp", ":8093")
//	if err != nil {
//		panic(err)
//	}
//
//	server := grpc.NewServer(grpc.ChainUnaryInterceptor(myInterceptor1, myInterceptor2))
//	proto.RegisterHelloServiceServer(server, &Server{})
//
//	if err := server.Serve(listener); err != nil {
//		panic(err)
//	}
//}
