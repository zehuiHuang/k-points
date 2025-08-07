package stack

import (
	"sort"
	"strconv"
	"strings"
)

//单调栈

/*
739. 每日温度：https://leetcode.cn/problems/daily-temperatures/description/
思路：使用单调栈，栈从上到下为单调递增(目的是为了获取当前i比最小的大多少，然后更新栈顶的差值，最后将栈顶弹出)
*/

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	//定义单调栈
	stack := []int{}

	for i := 0; i < length; i++ {
		//获取当前遍历的值,并和栈顶的值进行对比（栈里存储的是数组下标）
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			//获取栈顶的数据（温度的数组下标）
			index := stack[len(stack)-1]
			//弹出栈顶数据
			stack = stack[:len(stack)-1]
			ans[index] = i - index
		}
		stack = append(stack, i)
	}
	return ans
}

// 496. 下一个更大元素 I
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	mp := make(map[int]int)
	for i := range nums1 {
		mp[nums1[i]] = i
	}
	ans := []int{}
	for i := 0; i < len(nums1); i++ {
		ans = append(ans, -1)
	}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			peek := stack[len(stack)-1]
			if _, ok := mp[nums2[peek]]; ok {
				ans[mp[nums2[peek]]] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

/*
84. 柱状图中最大的矩形
思路：找到数组中左右两边分别第一个小于当前值的元素,比如：[2，1,5,6,2,3]中找到5的左右两边第一个小于他的值分别为1和2
*/
func largestRectangleArea(heights []int) int {
	ret := 0
	//单调栈
	stack := []int{}
	//将数组两边各加上0，保证每个数组元素一定能找到左右两边小于当前值的值
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			//弹出栈，找到栈中栈顶左边第一个小于栈顶的元素下标
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			ret = max(ret, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return ret
}

// 20. 有效的括号
func isValid(s string) bool {
	stack := []rune{}
	mp := make(map[rune]rune)
	mp[')'] = '('
	mp[']'] = '['
	mp['}'] = '{'
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, rune(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			}
			peek := stack[len(stack)-1]
			if peek != mp[rune(s[i])] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 1047. 删除字符串中的所有相邻重复项
func removeDuplicates(s string) string {
	stack := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := []string{}
	mp := make(map[string]bool)
	mp["+"] = true
	mp["-"] = true
	mp["*"] = true
	mp["/"] = true
	for i := 0; i < len(tokens); i++ {
		if !mp[tokens[i]] {
			stack = append(stack, tokens[i])
		} else {
			v1, _ := strconv.Atoi(stack[len(stack)-1])
			v2, _ := strconv.Atoi(stack[len(stack)-2])
			if tokens[i] == "*" {
				stack = stack[:len(stack)-2]
				res := v1 * v2
				stack = append(stack, strconv.Itoa(res))
			}
			if tokens[i] == "-" {
				stack = stack[:len(stack)-2]
				res := v2 - v1
				stack = append(stack, strconv.Itoa(res))
			}
			if tokens[i] == "+" {
				stack = stack[:len(stack)-2]
				res := v1 + v2
				stack = append(stack, strconv.Itoa(res))
			}
			if tokens[i] == "/" {
				stack = stack[:len(stack)-2]
				res := v2 / v1
				stack = append(stack, strconv.Itoa(res))
			}
		}
	}
	r, _ := strconv.Atoi(stack[0])
	return r
}

// 239. 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	stack := []int{}
	var push func(v int)
	//push下标
	push = func(index int) {
		for len(stack) > 0 && nums[index] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, nums[stack[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		if stack[0] < i-k+1 {
			stack = stack[1:]
		}
		ans = append(ans, nums[stack[0]])
	}
	return ans
}

// 347. 前 K 个高频元素
func topKFrequent(nums []int, k int) []int {
	ans := []int{}
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for key, _ := range mp {
		ans = append(ans, key)
	}

	sort.Slice(ans, func(i, j int) bool {
		return mp[ans[i]] > mp[ans[j]]
	})
	return ans[:k]
}

// 71. 简化路径
func simplifyPath(path string) string {
	stack := []string{}
	s := strings.Split(path, "/")
	for i := 0; i < len(s); i++ {
		if s[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s[i] != "" && s[i] != "." {
			stack = append(stack, s[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
