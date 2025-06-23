package base

import "testing"

// 引用传递
func Test_slice(t *testing.T) {
	s := []int{2, 3, 4}
	// [2,3,4] -> [-1,3,4]
	changeSlice(s)
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
	s1 := s[8:]
	s1[0] = -1
	t.Logf("s: %v", s)
}
