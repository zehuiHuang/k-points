package law

//特殊规律

/*
*最长回文子传
leetcode：5 https://leetcode-cn.com/problems/longest-palindromic-substring/

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

/*
*
 */
func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		print(left)
		print(right)
	}
	return left + 1, right - 1
}
