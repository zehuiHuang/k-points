package main

import (
	"fmt"
	"sync"
)

/*
*
三个 groutine 依次向 map 追加数据 0-30，要求协程依次执行。不限制资源数量，打印 执行时候的 协程编号、追加的数字，以及最后输出追加完成的map

示例：
go1: 0
go2: 1
go3: 2
go1: 3
......
[0:0 1:1 2:2 3:3 ..... 30:30]
*/
func main() {
	mp := make(map[int]int)
	//1,2,3
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	//数据准备
	//协程保证线程安全
	//g1
	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-c1:
				if !ok {
					return
				}
				if data == 30 {
					close(c1)
					close(c2)
					close(c3)
					return
				}
				mp[data] = data
				fmt.Printf("job:%s,value:%d\n", "1", data)
				c2 <- data + 1

			}
		}
	}()
	//g2
	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-c2:
				if !ok {
					return
				}
				if data == 30 {
					close(c1)
					close(c2)
					close(c3)
					return
				}
				mp[data] = data
				fmt.Printf("job:%d,value:%d\n", 2, data)
				c3 <- data + 1
			}
		}
	}()
	//g3
	go func() {
		defer wg.Done()
		for {
			select {
			case data, ok := <-c3:
				if !ok {
					return
				}
				if data == 30 {
					close(c1)
					close(c2)
					close(c3)
					return
				}
				mp[data] = data
				fmt.Printf("job:%d,value:%d\n", 3, data)
				c1 <- data + 1
			}
		}
	}()
	//启动
	c1 <- 0
	wg.Wait()
	fmt.Println(mp)
}
