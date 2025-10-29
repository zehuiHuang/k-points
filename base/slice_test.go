package base

import (
	"fmt"
	"testing"
)

// 引用传递
func Test_slice(t *testing.T) {
	s := []int{2, 3, 4}
	// [2,3,4] -> [-1,3,4]
	changeSlice(s)
	fmt.Print(s)
}

func Test_slice1(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice2(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice3(t *testing.T) {
	s := make([]int, 10, 11)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_Slice4(t *testing.T) {
	s := make([]int, 10, 12)
	//大小和容量为2和4
	s1 := s[8:]
	//s1的0就是s的8
	s1[0] = -1
	t.Logf("s: %v", s)
}

func Test_slice5(t *testing.T) {
	s := make([]int, 10, 12)
	//大小和容量为3和5,容量并不会被切分
	s1 := s[7:9]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

// 切片扩容后会新增一个slice header
func Test_slice6(t *testing.T) {
	s := make([]int, 10, 12)
	//大小和容量为2和4
	s1 := s[8:]
	s1 = append(s1, []int{10, 11, 12}...)
	v := s[10]
	fmt.Print(v)
	// ...
	// 此时会报错:数组越界,因为扩容会给s1新增一个slice header
}

func Test_slice7(t *testing.T) {
	s := make([]int, 10, 12)
	//大小和容量为2和4
	s1 := s[8:]
	//s1 append不会影响s,因为
	s1 = append(s1, []int{10, 11}...)
	fmt.Print(s)
	fmt.Print(s1)
	//变更s1的索引0位置会影响到s,变更s的9位置会影响到s1,因为s1的0位置就是s的8位置,s的9位置就是s1的1位置
	s1[0] = -1
	s[9] = -1
	fmt.Print(s)
	fmt.Print(s1)
}

// 重要!!! 切片进行方法参数传递时,实际传递的是一个副本,他们有相同的指针、长度和容量,
// 在内部进行append时,方法内的长度或容量变化后并不会影响原有切片的长度和容量
func Test_slice8(t *testing.T) {
	s := make([]int, 10, 12)
	//size:2 cap:4
	s1 := s[8:]
	changeSlice2(s1)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice2(s1 []int) {
	s1 = append(s1, 10)
	s1[0] = -1
}
