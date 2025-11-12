package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()

	//go func() {
	//	ch <- 1
	//}()
	for {
		select {
		case data, ok := <-ch:
			fmt.Printf("data:%v", data)
			fmt.Printf("ok:%v", ok)
			fmt.Println("rec----")
			fmt.Println(data)
		default:
			time.Sleep(1000)
			fmt.Println("default------")
		}
	}
	fmt.Println("11111")
}
