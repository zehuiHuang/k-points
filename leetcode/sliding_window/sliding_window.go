package sliding_window

import "math"

/**
滑动窗口 模版

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
			if cnt[k] < v { // 如果当前窗口中某个字符的数量小于需要的数量
				return false
			}
		}
		return true
	}

	// 使用滑动窗口遍历s
	for l, r := 0, 0; r < sLen; r++ {
		// 如果当前字符在t中出现过，增加其在窗口中的计数(只加存在t中的字符)
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

// 209:长度最小的子数组
// 2, 3, 1, 2, 4, 3
func minSubArrayLen(target int, nums []int) int {
	//滑动窗口
	n := len(nums)
	left := 0
	sum := 0
	result := n + 1
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			subLength := right - left + 1
			if result > subLength {
				result = subLength
			}
			sum -= nums[right]
			left++
		}
	}
	if result == n+1 {
		return 0
	} else {
		return result
	}
}

// 1004. 最大连续1的个数 III
/**
给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
示例 1：
输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6
*/
func longestOnes(nums []int, k int) int {
	//左指针，下标从0到left-1之间出现的0的数量，下标从0到right下标出现0的数量
	ans := 0
	left, lsum, rsum := 0, 0, 0
	for right, v := range nums {
		//统计另的个数
		rsum += 1 - v
		for rsum-lsum > k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 287. 寻找重复数
func findDuplicate(nums []int) int {
	//快慢指针
	slow, fast := 0, 0
	//先模拟环形节点,找到相遇的点
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	//将慢指针移动到头部,快慢指针同时移动一步,相遇的点即为公共节点
	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	return slow
}
