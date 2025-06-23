package b

import (
	"go.uber.org/dig"
	"testing"
)

func Test_dig(t *testing.T) {
	// 创建一个容器
	c := dig.New()

	// 注入各个 bean 的构造器函数
	_ = c.Provide(NewB)

	// 使用 bean A 的 struct 形式，与 container 进行 Invoke 交互
	var a A
	_ = c.Invoke(func(_a A) {
		a = _a
	})

	t.Logf("got a: %+v, got b: %+v", a, a.B)
}
