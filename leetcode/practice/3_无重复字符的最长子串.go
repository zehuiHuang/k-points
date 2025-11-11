package practice

// 1、滑动窗口 + 哈希表
func lengthOfLongestSubstring(s string) int {
	occ := make(map[rune]bool)
	left, ans := 0, 0
	//遍历字符串
	for right := 0; right < len(s); right++ {
		//判断当前字符串是否在occ中
		for occ[rune(s[right])] {
			//循环删除的目的是为了将left的位置移动到第一次出现重复字符的下一个位置，例如abcb，left需要移动到下标为1的b字节的下一个位置（即c）
			//若在，则删除该字符，并且左边界向右滑动一个位置
			delete(occ, rune(s[left]))
			left++
		}
		//给当前字符添加到occ中
		occ[rune(s[right])] = true
		//计算当前窗口的长度
		length := right - left + 1
		if length > ans {
			ans = length
		}
	}
	return ans
}

// 2、滑动窗口 + 哈希表
/**
"abcabcbb"

dp[j] 代表以字符 s[j] 为结尾的 “最长不重复子字符串” 的长度。

*/
func lengthOfLongestSubstring2(s string) int {
	return 0
}
