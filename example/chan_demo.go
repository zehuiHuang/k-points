package example

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

/*
*
1、读写未初始化的会引发死锁;
2、写已关闭的会发生panic;
3、读已关闭的会继续返回通道中的值,没有了就返回对应类型的零值;
4、重复close会报panic
*/
func chanExample1() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	// 安全：读已经关闭的 channel
	v1, ok1 := <-ch
	fmt.Println(v1, ok1) // 输出: 1 true

	v2, ok2 := <-ch
	fmt.Println(v2, ok2) // 输出: 2 true

	v3, ok3 := <-ch
	fmt.Println(v3, ok3) // 输出: 0 false（通道里没值了）

	// 危险操作：往已经关闭的 channel 写
	ch <- 3 // 这里会 panic
}

func chanExample12() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	time.Sleep(time.Duration(2) * time.Second)
	close(ch1)
}

func bb() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		ch1 <- 1
	}()
	<-ch1
	wg.Wait() // 等待发送完成

	close(ch1)
}

// 超时
func chanExample2() {
	reqResponded := make(chan bool, 1)
	timeout := time.Duration(10) * time.Second
	ctx := context.Background()
	fmt.Println("111111111")
	reqCtx := buildRetryCtx(ctx, timeout, reqResponded)
	fmt.Println("2222222")
	fmt.Println(reqCtx)
	highTime()
	fmt.Println("3333333")
	reqResponded <- true
	fmt.Println("4444444")
}

func highTime() {
	time.Sleep(time.Duration(12) * time.Second)
}

func buildRetryCtx(ctx context.Context, timeout time.Duration, reqResponded chan bool) context.Context {
	reqCtx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-reqResponded:
			fmt.Println("bbbbbbbb")
			return
		case <-time.After(timeout):
			fmt.Println("aaaaaaaa")
			cancel()
			return
		}
	}()
	return reqCtx
}

func SafeGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Sprintf("goroutine panic: %+v, stack: %s", r, string(debug.Stack()))
			}
		}()
		f()
	}()
}
