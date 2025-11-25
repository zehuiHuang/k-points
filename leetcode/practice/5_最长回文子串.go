package practice

/*
*
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
*/

// 思路:动态规划,通过每个字符串朝外扩散
// dp[i][j] 表示 s[i..j] 是否是回文串
/**
 */
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	res_len := 0
	res_start := 0
	for i := 0; i < len(s); i++ {
		//以单个字符未目标进行扩散
		L, R := i, i
		//起始位置
		for L >= 0 && R < len(s) && s[L] == s[R] {
			if res_len < R-L+1 {
				res_len = R - L + 1
				res_start = L
			}
			L--
			R++
		}
		//以双字符为目标进行扩撒
		L, R = i, i+1
		for L >= 0 && R < len(s) && s[L] == s[R] {
			if res_len < R-L+1 {
				res_len = R - L + 1
				res_start = L
			}
			L--
			R++
		}
	}
	return s[res_start : res_start+res_len]
}
