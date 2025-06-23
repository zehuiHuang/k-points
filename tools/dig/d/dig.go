package b

import "go.uber.org/dig"

type A struct {
	//dig.In 标识的方式替代构造函数，标志出A所有可导出的成员变量均为依赖项
	dig.In
	B1 *B `name:"b1"`
	B2 *B `name:"b2"`
}
type OutB struct {
	dig.Out
	//dig.Out 在provide 流程中将某个类的所有可导出成员属性均作为bean注入到container中
	// 分别提供名称为 b1 和 b2 的 bean
	B1 *B `name:"b1"`
	B2 *B `name:"b2"`
}

func NewOutB() OutB {
	return OutB{
		B1: NewB1(),
		B2: NewB2(),
	}
}

type B struct {
	Name string
}

func NewB1() *B {
	return &B{
		Name: "i am b111111",
	}
}

func NewB2() *B {
	return &B{
		Name: "i am b222222",
	}
}
