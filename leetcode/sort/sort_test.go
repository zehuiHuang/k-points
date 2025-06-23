package sort

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	arr := []int{1, 6, 3, 9, 0, 6, 3, 2}
	//fmt.Println(selectSort(arr))
	//fmt.Println(bubbleSort(arr))
	//fmt.Println(insertSort(arr))

	fmt.Println(quickSort(arr))
}

func Test_slice(t *testing.T) {
	s := make(map[int]int)
	changeSlice(s)
	t.Logf("s: %v", s)
}

func changeSlice(s1 map[int]int) {
	var s2 map[int]int
	s2 = s1
	s2[1] = 2
	s2[2] = 4
}

func Test_slice2(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice2(s1)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice2(s1 []int) {
	s1 = append(s1, 10)
}
