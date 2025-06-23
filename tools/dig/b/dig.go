package b

import (
	"go.uber.org/dig"
)

type A struct {
	dig.In
	B *B
}

type B struct {
	Name string
}

func NewB() *B {
	return &B{
		Name: "i am b",
	}
}
