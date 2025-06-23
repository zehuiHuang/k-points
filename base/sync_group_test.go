package base

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSyncGroup(t *testing.T) {
	taskNum := 10
	ch := make(chan struct{}, taskNum)

	for i := 0; i < taskNum; i++ {
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			//working
			<-time.After(time.Second)
		}()

	}

	for i := 0; i < taskNum; i++ {
		<-ch
	}
}

func TestNameCaptor1(t *testing.T) {
	taskNum := 10
	dataChan := make(chan interface{})
	resp := make([]interface{}, 0, taskNum)
	stopChan := make(chan struct{}, 1)

	go func() {
		for v := range dataChan {
			resp = append(resp, v)
		}
		stopChan <- struct{}{}
	}()

	var syncGroup sync.WaitGroup
	for i := 0; i < taskNum; i++ {
		syncGroup.Add(1)
		go func(ch chan<- interface{}) {
			defer syncGroup.Done()
			ch <- time.Now().UnixNano()
		}(dataChan)
	}
	syncGroup.Wait()
	close(dataChan)

	<-stopChan
	t.Logf("resp:%+v", resp)
}

func TestNameCaptor2(t *testing.T) {
	var counter uint64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.StoreUint64(&counter, counter+1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter = ", counter)
}
