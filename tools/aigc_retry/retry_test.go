package aigc_retry

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-learn/tools/aigc_retry/business"
	"go-learn/tools/aigc_retry/interceptor"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	var (
		retriableErrors = []ErrorCode{ErrorNeedRetry}
		retryTimeout    = 50 * time.Millisecond
	)
	var myInterceptor = func(ctx context.Context, req, reply any, invoker interceptor.UnaryInvoker) (err error) {
		// 添加前处理...
		fmt.Printf("myInterceptor preprocess, req: %+v\n", req)
		err = invoker(ctx, req, reply)
		// 添加后处理...
		fmt.Printf("myInterceptor postprocess, req: %+v\n", reply)
		return
	}
	var retryInterceptor = UnaryClientInterceptor(
		WithCodes(retriableErrors...),
		WithMax(3),
		WithBackoff(BackoffLinear(retryTimeout)),
	)
	var chainUnaryInts = []interceptor.UnaryClientInterceptor{myInterceptor, retryInterceptor}

	f := ChainUnaryInterceptors(chainUnaryInts)

	reply := ""
	req := ""
	f(context.Background(), req, &reply, func(ctx context.Context, req, reply any) error {
		reply = "SUCCESS"
		fmt.Printf("handler invoker-------------\n")
		return retriableErrors[0]
	})
}

//func Dispatch(c *gin.Context, taskInfo *business.Info, ret *bool) error {
//	//同步模式
//	result := true
//	err := apiDispatcher.Dispatch(c, taskInfo)
//	if err != nil {
//		close(taskInfo.RespChan)
//	}
//	ret = &result
//	taskInfo.ProductName = "hzhNameTest"
//	fmt.Printf("handler invoker-------------\n")
//	return nil
//}

var apiDispatcher *business.ApiDispatcher
var (
	retriableErrors = []ErrorCode{ErrorNeedRetry}
	retryTimeout    = 50 * time.Millisecond
)
var Dispatch = func(ctx context.Context, taskInfo *business.Info) (bool, error) {
	result := true
	taskInfo.ProductName = "hzhNameTest"
	err := apiDispatcher.Dispatch(ctx, taskInfo)
	if err != nil {
		close(taskInfo.RespChan)
	}
	fmt.Printf("handler invoker-------------\n")
	return result, retriableErrors[0]
}

func TestCreate(t *testing.T) {
	var myInterceptor = func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool, invoker interceptor.Invoker) (err error) {
		// 添加前处理...
		fmt.Printf("auth,limit,check preprocess\n")
		//TODO retry
		err = invoker(ctx, c, taskInfo, ret)
		// 添加后处理...
		fmt.Printf("auth,limit,check postprocesst\n")
		return
	}
	var retryInterceptor = AigcInterceptor(
		WithCodes(retriableErrors...),
		WithMax(3),
		WithBackoff(BackoffLinear(retryTimeout)),
		WithPerRetryTimeout(10*time.Second),
	)
	var c *gin.Context
	engine := New(c)

	engine.Use(Recovery())
	engine.Use(myInterceptor)
	engine.Use(retryInterceptor)

	var chainUnaryInts = []interceptor.HandlerFunc{myInterceptor, retryInterceptor}

	f := ChainInterceptors(chainUnaryInts)

	var ret bool
	taskInfo := business.Info{}
	err := f(context.Background(), c, &taskInfo, ret, func(ctx context.Context, c *gin.Context, taskInfo *business.Info, ret bool) error {
		ret, err := Dispatch(c, taskInfo)
		return err
	})
	fmt.Printf("handler invoker result error-----------------------------------------%v\n", err)
}

type node struct {
	handler interceptor.HandlerFunc
	next    []*node
}

type nodes struct {
	list []*node
}
