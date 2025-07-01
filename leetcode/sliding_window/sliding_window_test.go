package sliding_window

import "testing"

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
