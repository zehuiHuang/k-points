package sliding_window

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	//print(lengthOfLongestSubstring("abcabcbb"))
	//print(lengthOfLongestSubstring("abab"))
	//print(lengthOfLongestSubstring("abbbdf"))
	print(lengthOfLongestSubstring("abcb"))
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 2, 4, 3}
	minSubArrayLen(7, nums)
}

func TestLongestOnes(t *testing.T) {
	p := []int{1, 0, 1, 0, 1}
	fmt.Println(longestOnes(p, 1))
}
