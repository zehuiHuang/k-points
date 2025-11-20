package example

import (
	"fmt"
	"sync"
)

func poolExample() {
	// 1 模板类声明
	type instance struct {
		body []byte
	}

	// 2 对象池实例声明
	var pool sync.Pool

	// 3 对象池构造器函数声明
	pool.New = func() any {
		return &instance{
			body: []byte("{\"key\":\"value\"}"),
		}
	}

	// 4 应用对象池
	for i := 0; i < 100000; i++ {
		// 4.1 从对象池中获取对象实例
		inst, _ := pool.Get().(*instance)
		// 4.2 使用对象实例
		fmt.Printf("body: %s\n", inst.body)
		// 4.3 用完对象实例后，归还到对象池中
		pool.Put(inst)
	}
}

func b() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	close(ch1)
}

func a() {
	var wg sync.WaitGroup
	//定义通道
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	wg.Add(3)
	//1
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case data, ok := <-ch1:
				if !ok {
					return
				}
				fmt.Printf("job1:%v\n", data)
				ch2 <- data + 1
			}
		}
	}()
	//2
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case data, ok := <-ch2:
				if !ok {
					return
				}
				fmt.Printf("job2:%v\n", data)
				ch3 <- data + 1
			}
		}
	}()

	//3
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			select {
			case data, ok := <-ch3:
				if !ok {
					return
				}
				fmt.Printf("job3:%v\n", data)
				if i == 99 {
					close(ch1)
					close(ch2)
					close(ch3)
					return
				}
				ch1 <- data + 1
			}
		}
	}()
	ch1 <- 1
	wg.Wait()
}
