package example

import "fmt"

/*
*
1、读写未初始化的会引发死锁;
2、写已关闭的会发生panic;
3、读已关闭的会继续返回通道中的值,没有了就返回对应类型的零值;
4、重复close会报panic
*/
func chanExample1() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	// 安全：读已经关闭的 channel
	v1, ok1 := <-ch
	fmt.Println(v1, ok1) // 输出: 1 true

	v2, ok2 := <-ch
	fmt.Println(v2, ok2) // 输出: 2 true

	v3, ok3 := <-ch
	fmt.Println(v3, ok3) // 输出: 0 false（通道里没值了）

	// 危险操作：往已经关闭的 channel 写
	ch <- 3 // 这里会 panic
}
