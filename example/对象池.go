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
