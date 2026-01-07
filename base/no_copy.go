package base

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type noCopy struct {
}

func (*noCopy) Lock() {
}
func (*noCopy) UnLock() {
}

type copyChecker uintptr

func (c *copyChecker) check() {
	//1、检查当前存储在 c 中的值是否等于 c 自身的内存地址
	//如果相等，说明 c 未被复制且完成了初始化
	//如果不相等，说明 c 已被复制

	//2、原子操作进行初始化
	//如果成功则说明初始化成功
	//失败,则说明其他线程并发执行时失败,说明被copy了

	//3、二次比较,确保初始化成功

	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("sync.Cond is copied")
	}
}

// SafeCounter 禁止copy的结构体
type SafeCounter struct {
	// noCopy 禁止copy,会在运行时检查,依靠go vet 检查
	noCopy noCopy
	//运行时若发现直接panic
	checker copyChecker
}

// NewSafeCounter 保证单例模式
func NewSafeCounter() SafeCounter {
	sc := new(SafeCounter)
	fmt.Println(unsafe.Pointer(sc))
	return *sc
}

// 执行时候检查,若已经被copy则会直接panic
func (sc *SafeCounter) business() {
	sc.checker.check()
	fmt.Println("business")
}

func badCopy() {
	var sc1 SafeCounter
	var sc2 = sc1
	_ = sc2
}

func example() {
	//代码运行时未限制
	badCopy()

	//正常执行
	sc := NewSafeCounter()
	sc.business()

	//执行失败
	sc2 := sc
	sc2.business()
}
