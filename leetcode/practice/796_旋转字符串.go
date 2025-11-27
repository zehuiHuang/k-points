package practice

import "strings"

/**
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。
*/

// 思路:
// abcdeabcde
// cdeab

/*
*
思路:技巧类问题:,例如将s随便切一下氛围L和R ,即s=LR,将L移动到前面的到RL,两个RLRL 是一定包含s=LR
*/
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
