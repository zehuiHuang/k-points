package example

import "fmt"

/*
*
 */
type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() { fmt.Println("Woof") }

type Cat struct{}

func (c *Cat) Speak() { fmt.Println("Meow") }

func aff() {
	var s Speaker
	s = Dog{}
	s = &Dog{} //值接受者既可以是指针也可以是结构体
	//s = Cat{} //此时Cat不能复制给Speaker,只有指针类型的菜可以
	s = &Cat{}
	fmt.Println(s)
}
