package b

import (
	"go.uber.org/dig"
)

type A struct {
	dig.In
	// 依赖的 bean list
	Bs []*B `group:"b_group"`
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

type BGroup struct {
	dig.Out
	// 提供 bean list
	Bs []*B `group:"b_group,flatten"`
}

// 返回提供 bean list 的构造器函数
func NewBGroupFunc(bs ...*B) func() BGroup {
	return func() BGroup {
		group := BGroup{
			Bs: make([]*B, 0, len(bs)),
		}
		group.Bs = append(group.Bs, bs...)
		return group
	}
}
