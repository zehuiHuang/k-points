package b

import (
	"go.uber.org/dig"
	"testing"
)

func Test_dig(t *testing.T) {
	// 创建一个容器
	c := dig.New()

	// 注入各个 bean 的构造器函数
	_ = c.Provide(NewOutB)

	var a A
	_ = c.Invoke(func(_a A) {
		a = _a
	})

	t.Logf("got a: %+v, got b1: %+v, got b2: %+v", a, a.B1, a.B2)
}
