package method

import (
	"fmt"
)

/**
方法的本质就是函数，只是
T.F1(method.T{Name: "eggo"})
*/

type T struct {
	Name string
}

func (t T) F1() {
	fmt.Println(t.Name)
}

/*
不同点：MyType1相当于给int32起的别名，属于同一个类型，类似rune类型；而MyType2相当于创建了一个新类型，有自己的类型元数据
*/

type MyType1 = int32
type MyType2 int32

var a rune
