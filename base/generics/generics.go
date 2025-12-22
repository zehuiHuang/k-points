package main

import "fmt"

//泛型

// 定义接口类型
type iface interface {
	Foo()
}

//定义泛幸接口类型,Get返回的是实现了iface接口的结构体

type iface2[T iface] interface {
	Get() T
}

// 定义结构体,实现iface接口
type s struct {
}

// 定义结构体
type s2 struct {
	S *s
}

func (s *s) Foo() {
	fmt.Println("------- foo ------------")
}

func (s2 *s2) Get() *s {
	return s2.S
}

// 调用时,必须是iface接口实现的结构体
func test[T iface](obj iface2[T]) {
	obj.Get().Foo()
}
func NewS2() *s2 {
	return &s2{
		S: &s{},
	}
}
