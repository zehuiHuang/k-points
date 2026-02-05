package main

import (
	"sync"
)

// 通过对象池技术优化切片分配性能
var uint64sPool sync.Pool

type Uint64s struct {
	A []uint64
}

func GetUint64s(size int) *Uint64s {
	v := uint64sPool.Get()
	if v == nil {
		return &Uint64s{
			A: make([]uint64, size),
		}
	}
	is := v.(*Uint64s)
	// 关键技巧：复用底层数组，仅调整切片长度
	// 避免了重新 make([]uint64) 的开销
	is.A = SetLength(is.A, size)
	return is
}

func SetLength[T any](a []T, newLen int) []T {
	// 如果新长度大于当前容量，则扩容
	if n := newLen - cap(a); n > 0 {
		a = append(a[:cap(a)], make([]T, n)...)
	}
	return a[:newLen]
}
