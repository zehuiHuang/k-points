package double_pointer

import (
	"fmt"
	"testing"
)

func TestAbc(t *testing.T) {
	s := "Hello, 世界！"

	// 错误方式：按字节遍历（会乱码）
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i]) // 输出：Hello, ä¸–ç•Œï¼
	}

	// 正确方式：按 rune 遍历
	for _, r := range s {
		fmt.Printf("%c", r) // 输出：Hello, 世界！
	}
}

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	findAnagrams(s, p)
}
