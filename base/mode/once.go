package mode

import (
	"fmt"
	"sync"
)

// 单例模式
var (
	s    *singletonInstance
	Once sync.Once
)

type singletonInstance struct {
}

type Instance interface {
	DoWork()
}

func (s singletonInstance) DoWork() {
	fmt.Println("1111")
}

func newSingleInstance() *singletonInstance {
	return &singletonInstance{}
}

func GetInstance() Instance {
	Once.Do(func() {
		s = newSingleInstance()
	})
	return s
}
