package main

import "runtime"

// 内存中数据块的结构体

type inmemoryPart struct {
	creationTime uint64
}

func (mp *inmemoryPart) Reset() {
	mp.creationTime = 0
}

// 容量严格限制为 CPU 核数，防止内存无限膨胀
// cgroup.AvailableCPUs()
var mpPool = make(chan *inmemoryPart, runtime.NumCPU())

// getInmemoryPart 从对象池中获取 inmemoryPart 实例
func getInmemoryPart() *inmemoryPart {
	select {
	case mp := <-mpPool: // 尝试从池中获取
		return mp
	default:
		return &inmemoryPart{} // 池空了，才新建
	}
}

// putInmemoryPart 归还 inmemoryPart 实例到对象池
func putInmemoryPart(mp *inmemoryPart) {
	mp.Reset()
	select {
	case mpPool <- mp: // 尝试归还
	default:
		// 池满了，直接丢弃，等待 GC 回收
	}
}
