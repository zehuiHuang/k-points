package law

import (
	"strconv"
	"strings"
)

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

// 415. 字符串相加
/**
输入：num1 = "456", num2 = "77"
输出："533"
*/
// 思路:双指针,进=进位、字符转数字和数字转字符串
func addStrings(num1 string, num2 string) string {
	//滑动窗口
	res := ""
	//进位
	carry := 0
	w1 := len(num1) - 1
	w2 := len(num2) - 1
	for w1 >= 0 || w2 >= 0 {
		n1, n2 := 0, 0
		if w1 >= 0 {
			n1 = int(num1[w1] - '0')
		}
		if w2 >= 0 {
			n2 = int(num2[w2] - '0')
		}
		tmp := n1 + n2 + carry
		carry = tmp / 10
		v := tmp % 10
		res = strconv.Itoa(v) + res
		w1--
		w2--
	}
	if carry != 0 {
		return strconv.Itoa(carry) + res
	}
	return res
}

// 796. 旋转字符串
// 思路:双指针,一个放左边,一个放右边
// abcdeabcde
// cdeab
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	c := goal + goal
	if strings.Contains(c, s) {
		return true
	}
	return false
}

// 476. 数字的补数
func findComplement(num int) int {
	//思路:获取num二进制对应的mark,然后对num二进制取反,并与mark取&做位运算
	mark := 0
	v := num
	for v > 0 {
		//将mark左移动一位,并在末尾换成1
		mark = (mark << 1) | 1
		//被处理的数右移一位,知道被减少到0为止
		v >>= 1
	}
	//&mark的目的是为了移除掉高位的0
	return ^num & mark
}
