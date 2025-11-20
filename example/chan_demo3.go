package example

import (
	"fmt"
	"time"
)

/*
*通过信号优雅停机
定时任务定时处理
*/
func server(quit <-chan bool) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer func() {
		ticker.Stop()
		fmt.Println("stop ticker")
	}()

	for {
		select {
		case <-ticker.C:
			// 定期执行的任务
			fmt.Println("Processing...")
		case <-quit:
			// 收到退出信号
			fmt.Println("Server shutting down...")
			// 清理资源
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Server stopped")
			return
		}
	}
}

func chanDemo3Example() {
	quit := make(chan bool)
	go server(quit)

	time.Sleep(3 * time.Second)
	quit <- true // 发送退出信号

	time.Sleep(500 * time.Millisecond)
	fmt.Println("系统退出")

}
