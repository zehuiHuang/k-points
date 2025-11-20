package example

import (
	"fmt"
	"time"
)

/*
*
通道的case获取是随机的
*/
func basicSelect2() {
	go func() {
		ch1 := make(chan string, 1)
		ch2 := make(chan string, 1)

		ch1 <- "hello" // ch1 现在有数据可读
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg) // 这个会被执行
		case ch2 <- "world":
			fmt.Println("Sent to ch2")
		default:
			fmt.Println("Default case")
		}
	}()
	time.Sleep(time.Duration(15) * time.Second)
}

func basicSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	//ch := make(chan string, 1)

	ch1 <- "hello"

	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg) // 这个会被执行
	case ch2 <- "world":
		fmt.Println("Sent to ch2")
	default:
		fmt.Println("Default case")
	}

	//done := make(chan bool)
	//
	//for {
	//	select {
	//	case data := <-ch:
	//		fmt.Println("Received:", data)
	//	case <-time.After(time.Second):
	//		fmt.Println("Timeout!")
	//	case <-done:
	//		return // 退出循环
	//	}
	//}
}
