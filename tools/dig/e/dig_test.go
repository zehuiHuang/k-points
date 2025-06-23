package b

import (
	"go.uber.org/dig"
	"testing"
)

func Test_dig(t *testing.T) {
	// 创建一个容器
	c := dig.New()

	// 注入各个 bean 的构造器函数
	_ = c.Provide(NewBGroupFunc(NewB1(), NewB2()))

	var a A
	_ = c.Invoke(func(_a A) {
		a = _a
	})

	t.Logf("got a: %+v, got b1: %+v, got b2: %+v", a, a.Bs[0], a.Bs[1])
}
