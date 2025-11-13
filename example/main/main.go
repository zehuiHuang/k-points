package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1

	//go func() {
	//	ch <- 1
	//}()
	//for {
	select {
	case data, ok := <-ch:
		fmt.Printf("data:%v\n", data)
		fmt.Printf("ok:%v", ok)
		fmt.Println("rec----")
	default:
		time.Sleep(1000)
		fmt.Println("default------")
	}
	//}
	time.Sleep(time.Second * 3)
	fmt.Println("11111")
}

/**
要求实现一个 map：

（1）面向高并发；

（2）只存在插入和查询操作 O(1)；

（3）查询时，若 key 存在，直接返回 val；若 key 不存在，阻塞直到 key val 对被放入后，获取 val 返回； 等待指定时长仍未放入，返回超时错误；

（4）写出真实代码，不能有死锁或者 panic 风险.
*/

type MyConcurrentMap struct {
	MP   map[int]int
	Ch   chan struct{}
	lock *sync.RWMutex
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.lock.Lock()
	defer m.lock.Unlock()
	close(m.Ch)
	m.MP[k] = v
	m.Ch = make(chan struct{})
}

func (m *MyConcurrentMap) Get(k int, maxWaitingDuration time.Duration) (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if value, ok := m.MP[k]; ok {
		return value, nil
	}
	timeoutCh := time.After(maxWaitingDuration)
	for {
		select {
		case <-m.Ch:
			if value, ok := m.MP[k]; ok {
				return value, nil
			}
		case <-timeoutCh:
			return 0, errors.New("xx")
		}
	}
}
