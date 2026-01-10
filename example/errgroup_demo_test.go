package example

import "testing"

// 错误返回
func TestErrGroupExample(t *testing.T) {
	errGroupExample()
}

// 上下文取消
func TestErrGroup2Example(t *testing.T) {
	errGroupExample2()
}

// 限制并发数
func TestErrGroup3Example(t *testing.T) {
	errGroupExample3()
}

// 尝试启动
func TestErrGroup4Example(t *testing.T) {
	errGroupExample4()
}
