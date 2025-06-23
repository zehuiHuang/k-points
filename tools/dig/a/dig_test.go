package a

import (
	"go.uber.org/dig"
	"testing"
)

// 使用示例
func Test_dig(t *testing.T) {
	// 创建一个容器
	c := dig.New()

	// 注入各个 bean 的构造器函数
	_ = c.Provide(NewB)
	_ = c.Provide(NewA)

	// 注入 bean 获取器函数，并通过闭包的方式从中取出 bean
	var a *A
	_ = c.Invoke(func(_a *A) {
		a = _a
	})
	t.Logf("got a: %+v, got b: %+v", a, a.b)
}
