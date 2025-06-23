package b

import "go.uber.org/dig"

type A struct {
	dig.In
	B *B
	C *C
}

type B struct {
	Name string
}

func NewB() *B {
	return &B{
		Name: "i am b",
	}
}

type C struct {
	Age int
}

func NewC() *C {
	return &C{
		Age: 10,
	}
}

// 内置了 dig.Out
type OutBC struct {
	dig.Out
	B *B
	C *C
}

// 返回 struct 类型，不得使用 pointer
func NewOutBC() OutBC {
	return OutBC{
		B: NewB(),
		C: NewC(),
	}
}
