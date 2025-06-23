package base

import (
	"unsafe"
)

/**
1、切片的传递是引用传递
*/

// 1、数据结构
type slice struct {
	// 指向起点的地址
	array unsafe.Pointer
	// 切片长度
	len int
	// 切片容量
	cap int
}

// 2、初始化
/**
var s []int
b := make([]int,8)
s := make([]int,8,16)
s := []int{2,3,4}
*/

// 3、引用传递
func changeSlice(s []int) {
	s[0] = -1
}

//4、内容截取
/**
s := []int{1, 2, 3, 4, 5}
	// s1: [2,3,4,5]
	s1 := s[1:]
	// s2: [1,2,3,4]
	s2 := s[:len(s)-1]
	// s3: [2,3,4]
	s3 := s[1 : len(s)-1]
*/
