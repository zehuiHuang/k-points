package a

// bean A，内部又依赖了 bean B
type A struct {
	b *B
}

// bean A 构造器函数
func NewA(b *B) *A {
	return &A{
		b: b,
	}
}

// bean B
type B struct {
	Name string
}

// bean B 构造器函数
func NewB() *B {
	return &B{
		Name: "i am b",
	}
}
