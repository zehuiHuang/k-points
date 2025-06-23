package b

import (
	"go.uber.org/dig"
	"testing"
)

func Test_dig(t *testing.T) {
	// 创建一个容器
	c := dig.New()

	// 注入各 dig.Out 的构造器函数，需要是 struct 类型
	_ = c.Provide(NewOutBC)

	var a A
	_ = c.Invoke(func(_a A) {
		a = _a
	})

	t.Logf("got a: %+v, got b: %+v, got c: %+v", a, a.B, a.C)
}
