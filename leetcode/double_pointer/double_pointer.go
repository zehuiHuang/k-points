package double_pointer

import (
	"sort"
	"strings"
)

/*
*
leetcode 11. 盛最多水的容器
思路：双指针，左右指针分别指向数组的两端，计算当前容器的容量，然后移动高度较小的指针，直到两个指针相遇
*/
func maxArea(height []int) int {
	//左右指针
	left, right := 0, len(height)
	maxArea := 0
	//只要左指针小于右指针，则进行循环
	for left < right {
		//当前的容量
		curArea := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, curArea)
		//尝试移动时，要移动较小的指针，因为较小的指针决定了当前的容量
		if height[left] < height[right] {
			left = left + 1
		} else {
			right--
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

/*
*
leetcode 283. 移动零
示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
过程
1、0,1,0,3,12
2、1,0,0,3,12
3、1,3,0,0,12
4、1,3,12,0,0

思路：双指针，找到右追针不为零的值，并和左指针交换，然后left++（保证左指针的左边一定不为0），最后的结果肯定是左指针的后面都是0
*/
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	//只要右指针小于数组长度，则进行循环
	for right < n {
		//如果右指针不为0，则交换左右指针的值
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		//右指针右移
		right++
	}
}

/*
*
leetcode:15 三数之和

输入：nums = [-1,0,1,2,-1,-4]
sort: [-4,-1,-1,0,1,2]
输出：[[-1,-1,2],[-1,0,1]]
思路：双指针：
*/
func threeSum(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		first := nums[i]
		//如果第一个数大于0，则直接退出
		if first > 0 {
			break
		}
		target := 0 - first
		left := i + 1
		right := len(nums) - 1
		for left < right {
			twoSum := nums[left] + nums[right]
			if twoSum < target {
				left++
			} else if twoSum > target {
				right--
			} else {
				second := nums[left]
				third := nums[right]
				v := []int{first, second, third}
				ans = append(ans, v)
				//对第二个数去掉重复的
				for left < right && nums[left] == second {
					left++
				}
				//对第三个数去掉重复的
				for left < right && nums[right] == third {
					right--
				}
			}
		}
		//对第一个数去掉重复的
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

/*
leetcode42:接雨水
*/
//双指针
/**
理解:
1、左指针和右指针分别在数组的最两侧
左指针的接雨水计算：
要求2点：
1）左指针位置的高度比右指针矮（只有这样才能保证左指针这边的位置一定能接住水）
2）左指针维护一个最大高度的变量，表示左指针位置的最高高度，leftMax - height[left]即为当前左指针位置的接雨水量

右指针的接雨水计算与左指针类似...
*/
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right { // 双指针向中间靠拢
		leftMax = max(leftMax, height[left])    // 更新左侧最大高度
		rightMax = max(rightMax, height[right]) // 更新右侧最大高度
		if height[left] < height[right] {       // 较矮侧为左指针
			ans += leftMax - height[left] // 累加当前左位置的雨水量
			left++                        // 左指针右移
		} else { // 较矮侧为右指针
			ans += rightMax - height[right] // 累加当前右位置的雨水量
			right--                         // 右指针左移
		}
	}
	return
}

/*
*
leetcode：438. 找到字符串中所有字母异位词
思路：活动窗口，以数组作为窗口，并记录数组中的字符出现的次数，然后开始滑动
*/
func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	//建立两个数组存放字符串中字母出现的词频，并以此作为标准比较
	var sCount, pCount [26]int
	//当滑动窗口的首位在s[0]处时 （相当于放置滑动窗口进入数组）
	for i, ch := range p {
		sCount[s[i]-'a']++ //记录s中前pLen个字母的词频
		pCount[ch-'a']++   //记录要寻找的字符串中每个字母的词频(只用进行一次来确定)
	}
	//判断放置处是否有异位词     (在放置时只需判断一次)
	if sCount == pCount {
		ans = append(ans, 0)
	}

	//开始让窗口进行滑动
	for i, ch := range s[:sLen-pLen] { //i是滑动前的首位
		sCount[ch-'a']--        //将滑动前首位的词频删去
		sCount[s[i+pLen]-'a']++ //增加滑动后最后一位的词频（以此达到滑动的效果）

		//判断滑动后处，是否有异位词
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}

// 151. 反转字符串中的单词
/**
思路:
1、将指针放到末尾,然后向左滑动,只要不等于' '就累计凭借单词, 如果字符等于' ',则需要将拼接的单词放入集合,
2、当i=0时,将最后一个拼接的单词放入集合即可
*/
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	res := []string{}
	i := len(s) - 1
	//排除掉两边的空格
	//s = strings.Trim(s, " ")

	tmp := ""
	for i >= 0 {
		if i == 0 {
			tmp = string(s[i]) + tmp
			res = append(res, tmp)
			break
		}
		if s[i] != ' ' {
			tmp = string(s[i]) + tmp
		} else {
			if tmp != "" {
				res = append(res, tmp)
				tmp = ""
			}
			tmp = ""
		}
		i--
	}
	return strings.Join(res, " ")
}
