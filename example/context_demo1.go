package example

import (
	"context"
	"fmt"
	"time"
)

func contextDemo1Example() {
	// 创建一个子节点的context,3秒后自动取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))

	go watch(ctx, "监控1")
	go watch(ctx, "监控2")

	fmt.Println("现在开始等待5秒,time=", time.Now().Unix())
	time.Sleep(5 * time.Second)

	fmt.Println("等待5秒结束,准备调用cancel()函数，发现两个子协程已经结束了，time=", time.Now().Unix())
	cancel()
}

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}
