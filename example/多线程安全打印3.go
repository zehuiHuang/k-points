package example

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// syncMap 线程安全
func benchmarkSyncMap() {
	numThreads := 100
	numOps := 1000000

	var m sync.Map
	var opsCount uint64
	var wg sync.WaitGroup
	wg.Add(numThreads)

	startTime := time.Now()

	for i := 0; i < numThreads; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				//验证现场安全
				m.Store(j, j)
				//验证atomic的线程安全
				atomic.AddUint64(&opsCount, 1)
			}
		}()
	}

	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Printf("SyncMap: %d ops in %s\n", atomic.LoadUint64(&opsCount), elapsed)
}

func accessMap(m map[int]int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		m[i] = i
	}
}

// Map 线程不安全
func benchmarkMap() {
	numThreads := 100
	numOps := 1000000

	m := make(map[int]int)
	var wg sync.WaitGroup
	wg.Add(numThreads)

	startTime := time.Now()

	for i := 0; i < numThreads; i++ {
		go accessMap(m, numOps, &wg)
	}

	wg.Wait()

	elapsed := time.Since(startTime)

	fmt.Printf("Map: %d ops in %s\n", numThreads*numOps, elapsed)
}
