package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// 实现本地化 Worker Pool 的方案，它们都旨在减少 CPU 间的通信开销，提高缓存利用率
// 此处采用了绑定本地的P的方式
// PLocalWorkerPool 基于 Go 调度器 P 的本地工作池
type PLocalWorkerPool struct {
	numProcs   int
	localPools []*pLocalPool
	stopChan   chan struct{}
	wg         sync.WaitGroup
}

type pLocalPool struct {
	queue chan func()
	// 本地统计
	processed uint64
	// 使用 sync.Pool 减少内存分配
	taskPool sync.Pool
}

// NewPLocalWorkerPool 创建基于 P 的本地工作池
func NewPLocalWorkerPool() *PLocalWorkerPool {
	numProcs := runtime.GOMAXPROCS(0)

	pool := &PLocalWorkerPool{
		numProcs:   numProcs,
		localPools: make([]*pLocalPool, numProcs),
		stopChan:   make(chan struct{}),
	}

	// 初始化每个 P 的本地池
	for i := 0; i < numProcs; i++ {
		localPool := &pLocalPool{
			queue: make(chan func(), 256), // 较大的缓冲区
			taskPool: sync.Pool{
				New: func() interface{} {
					return new(taskWrapper)
				},
			},
		}
		pool.localPools[i] = localPool
	}

	// 启动 worker
	for i := 0; i < numProcs; i++ {
		pool.startWorker(i)
	}

	return pool
}

type taskWrapper struct {
	fn func()
}

// startWorker 启动 worker
func (p *PLocalWorkerPool) startWorker(pID int) {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		localPool := p.localPools[pID]

		for {
			select {
			case task := <-localPool.queue:
				if task != nil {
					task()
					atomic.AddUint64(&localPool.processed, 1)
				}
			case <-p.stopChan:
				return
			}
		}
	}()
}

// Submit 提交任务到当前 P 的本地队列
func (p *PLocalWorkerPool) Submit(task func()) {
	if task == nil {
		return
	}

	// 获取当前 goroutine 关联的 P
	// 注意：Go 没有直接获取当前 P ID 的公共 API
	// 这里使用一种近似方法：通过 goroutine ID 的哈希值
	pID := p.getCurrentPID()

	select {
	case p.localPools[pID].queue <- task:
	default:
		// 如果当前 P 的队列满了，尝试其他队列
		for i := 0; i < p.numProcs; i++ {
			select {
			case p.localPools[(pID+i)%p.numProcs].queue <- task:
				return
			default:
				continue
			}
		}
		// 如果所有队列都满了，同步执行
		task()
	}
}

// getCurrentPID 获取当前 P 的 ID（近似实现）
func (p *PLocalWorkerPool) getCurrentPID() int {
	// 在实际应用中，可以使用更精确的方法，如：
	// 1. 通过 runtime 的私有 API（不推荐）
	// 2. 使用线程本地存储（TLS）
	// 3. 使用 goroutine 特定的 ID

	// 这里使用简单的哈希方法作为示例
	return int(uintptr(goid()) % uintptr(p.numProcs))
}

// 模拟获取 goroutine ID（实际项目中不要使用）
func goid() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 解析栈跟踪中的 goroutine ID
	// 这里只是示例，实际需要解析
	return uint64(n)
}

// Stop 停止工作池
func (p *PLocalWorkerPool) Stop() {
	close(p.stopChan)
	p.wg.Wait()

	for _, pool := range p.localPools {
		close(pool.queue)
	}
}
