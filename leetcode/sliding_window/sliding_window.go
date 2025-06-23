package sliding_window

import "math"

/**
滑动窗口 模版

*/

/**
Set<Character> occ = new HashSet<Character>();
        int left = 0, ans = 0;
        for(int i = 0; i < s.length(); i++){
            while(occ.contains(s.charAt(i))){
                occ.remove(s.charAt(left++));
            }
            occ.add(s.charAt(i));
            ans = Math.max(ans, i - left + 1);
        }
        return ans;
*/

/*
*leetcode 3.无重复字符的最长子串
案例：
abcabcbb
abcb
*/
func lengthOfLongestSubstring(s string) int {
	occ := make(map[rune]bool)
	left, ans := 0, 0
	//遍历字符串
	for i := 0; i < len(s); i++ {
		//判断当前字符串是否在occ中
		for occ[rune(s[i])] {
			//循环删除的目的是为了将left的位置移动到第一次出现重复字符的下一个位置，例如abcb，left需要移动到下标为1的b字节的下一个位置（即c）
			//若在，则删除该字符，并且左边界向右滑动一个位置
			delete(occ, rune(s[left]))
			left++
		}
		//给当前字符添加到occ中
		occ[rune(s[i])] = true
		//计算当前窗口的长度
		length := i - left + 1
		if length > ans {
			ans = length
		}
	}
	return ans
}

/*
*
leetcode:76. 最小覆盖子串
描述：给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
例如：输入： ADOBECODEBANC  ABC
输出 BANC

思路：滑动窗口
1、统计t中每个字符的出现次数
2、遍历s时，累积每个字符出现的次数，同时检测累计后的字符串是否包含t中所有字符，若包含，则记录当前窗口的长度，并尝试缩小窗口，直到不包含t中所有字符为止
*/
func minWindow(s string, t string) string {
	// 初始化两个map：
	// ori: 记录t中每个字符需要的数量
	// cnt: 记录当前窗口中每个字符的数量
	ori, cnt := map[rune]int{}, map[rune]int{}

	// 统计t中每个字符的出现次数
	for i := 0; i < len(t); i++ {
		ori[rune(t[i])]++
	}

	sLen := len(s)
	len := math.MaxInt32 // 记录最小窗口长度
	ansL, ansR := -1, -1 // 记录最小窗口的左右边界

	// check函数：检查当前窗口是否包含t的所有字符
	check := func() bool {
		for k, v := range ori {
			if cnt[rune(k)] < v { // 如果当前窗口中某个字符的数量小于需要的数量
				return false
			}
		}
		return true
	}

	// 使用滑动窗口遍历s
	for l, r := 0, 0; r < sLen; r++ {
		// 如果当前字符在t中出现过，增加其在窗口中的计数
		if r < sLen && ori[rune(s[r])] > 0 {
			cnt[rune(s[r])]++
		}

		// 当窗口包含t的所有字符时，尝试缩小窗口
		for check() && l <= r {
			// 更新最小窗口
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}

			// 如果左边界字符在t中，减少其计数
			if _, ok := ori[rune(s[l])]; ok {
				cnt[rune(s[l])] -= 1
			}
			l++ // 左边界右移
		}
	}

	// 如果没有找到符合条件的子串，返回空字符串
	if ansL == -1 {
		return ""
	}
	// 返回最小覆盖子串
	return s[ansL:ansR]
}
