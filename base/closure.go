package base

import "fmt"

/**
https://mp.weixin.qq.com/s/iFYkcLbNK5pOA37N7ToJ5Q
https://zhuanlan.zhihu.com/p/645853924
*/
/*
*闭包理解和定义：
1、闭包最常见的方式是引用了其外层函数定义的局部变量，并以函数的方式返回
2、闭包也称为有状态的函数:各自持有自己的捕获列表
3、闭包只是拥有一个或多个捕获变量的Function Value而已
*/
func addGenerator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func doLater(msg string) func() {
	return func() {
		fmt.Println("Later:", msg)
	}
}

func doLater2(person *person) func() {
	return func() {
		fmt.Println("Later:", person.name)
	}
}

type person struct {
	name string
	age  int
}

func forEach(numbers []int, callback func(int)) {
	for _, num := range numbers {
		callback(num)
	}
}

func newCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

/**
总结
（1）Go语言里Function Value本质上是指向funcval结构体的指针；
（2）Go语言里闭包只是拥有捕获列表的Function Value；
（3）捕获变量在外层函数与闭包函数中要保持一致。
*/
