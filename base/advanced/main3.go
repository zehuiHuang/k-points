package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//此处有一个正确模糊的知识点:就是只要关闭了一个chan,那么所有在接收该chan的goroutine都会收到一个零值，并且ok==false

//实现本地化 Worker Pool 的方案，它们都旨在减少 CPU 间的通信开销，提高缓存利用率

// LocalWorkerPool 本地化工作池
type LocalWorkerPool struct {
	numWorkers int
	// 每个 worker 有自己的任务队列
	taskQueues []chan func()
	// 每个 worker 的统计信息
	stats []workerStat
	// 轮询索引，用于分发任务
	nextWorker uint64
	wg         sync.WaitGroup
	stopChan   chan struct{}
}

type workerStat struct {
	processed uint64
	queueLen  int
}

// NewLocalWorkerPool 创建本地化工作池
func NewLocalWorkerPool(numWorkers int) *LocalWorkerPool {
	if numWorkers <= 0 {
		numWorkers = runtime.NumCPU()
	}

	pool := &LocalWorkerPool{
		numWorkers: numWorkers,
		taskQueues: make([]chan func(), numWorkers),
		stats:      make([]workerStat, numWorkers),
		stopChan:   make(chan struct{}),
	}

	// 为每个 worker 创建缓冲通道
	for i := 0; i < numWorkers; i++ {
		// 设置合理的缓冲区大小
		pool.taskQueues[i] = make(chan func(), 64)
	}

	// 启动 worker
	for i := 0; i < numWorkers; i++ {
		pool.startWorker(i)
	}

	return pool
}

// startWorker 启动一个 worker goroutine
func (p *LocalWorkerPool) startWorker(id int) {
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()

		// 尝试将 goroutine 绑定到特定的 CPU 核心
		runtime.LockOSThread()
		// 在实际生产环境中，可以使用 syscall.SetAffinity 来设置 CPU 亲和性

		for {
			select {
			case task := <-p.taskQueues[id]:
				if task != nil {
					task()
					atomic.AddUint64(&p.stats[id].processed, 1)
				}
			case <-p.stopChan:
				return
			}

			// 更新队列长度统计
			p.stats[id].queueLen = len(p.taskQueues[id])
		}
	}()
}

// Submit 提交任务（使用轮询策略）
func (p *LocalWorkerPool) Submit(task func()) {
	if task == nil {
		return
	}

	// 使用原子操作获取下一个 worker 索引
	idx := atomic.AddUint64(&p.nextWorker, 1) % uint64(p.numWorkers)

	select {
	case p.taskQueues[idx] <- task:
		// 任务成功提交
	default:
		// 如果队列已满，尝试其他队列
		for i := 0; i < p.numWorkers; i++ {
			select {
			case p.taskQueues[(idx+uint64(i))%uint64(p.numWorkers)] <- task:
				return
			default:
				continue
			}
		}
		// 如果所有队列都满了，同步执行
		task()
	}
}

// SubmitTo 将任务提交到特定的 worker
func (p *LocalWorkerPool) SubmitTo(workerID int, task func()) error {
	if workerID < 0 || workerID >= p.numWorkers {
		return fmt.Errorf("invalid worker ID")
	}

	select {
	case p.taskQueues[workerID] <- task:
		return nil
	default:
		// 如果队列已满，直接执行
		task()
		return nil
	}
}

// Stop 停止工作池
func (p *LocalWorkerPool) Stop() {
	close(p.stopChan)
	p.wg.Wait()

	// 关闭所有任务队列
	for _, q := range p.taskQueues {
		close(q)
	}
}

// GetStats 获取统计信息
func (p *LocalWorkerPool) GetStats() []workerStat {
	stats := make([]workerStat, p.numWorkers)
	for i := 0; i < p.numWorkers; i++ {
		stats[i] = workerStat{
			processed: atomic.LoadUint64(&p.stats[i].processed),
			queueLen:  p.stats[i].queueLen,
		}
	}
	return stats
}
