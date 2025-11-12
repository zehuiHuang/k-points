package example

import (
	"sync"
)

/*
*
分别打印dog1、dog2、dog3各100次,且是安装顺序的,比如:
dog1
dog2
dog3
dog1
dog2
dog3
...
利用chan 的安全性,按照顺序并发打印日志
*/
func example1() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{}) // 控制 dog1 的打印
	ch2 := make(chan struct{}) // 控制 dog2 的打印
	ch3 := make(chan struct{}) // 控制 dog3 的打印
	wg.Add(3)

	// dog1 打印协程
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch1 // 等待触发信号
			println("dog1")
			ch2 <- struct{}{} // 触发 dog2
		}
	}()
	// dog2 打印协程
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch2 // 等待触发信号
			println("dog2")
			ch3 <- struct{}{} // 触发 dog3
		}
	}()

	// dog3 打印协程
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch3 // 等待触发信号
			println("dog3")
			if i < 99 { // 最后一轮不发送
				ch1 <- struct{}{} // 触发下一轮 dog1
			}
		}
	}()

	// 启动第一轮打印
	ch1 <- struct{}{}

	// 等待所有协程完成
	wg.Wait()
}
