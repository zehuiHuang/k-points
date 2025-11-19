package practice

import "sort"

// 思路:滑动窗口,找到其中一个值作为target,然后从左右两边判定,根据大小情况进行左滑动或右滑动
// 重点:注意重复数,如果left为v,但其右边还有值为v的,那么需要去重(左指针向右滑动),还有right值为v,right的左边还有v值,同理
// 第三个重复是选择的target右边有重复,则也需要去重
func threeSum(nums []int) [][]int {
	//排序
	sort.Ints(nums)
	ans := make([][]int, 0)
	//双指针
	for i := 0; i < len(nums)-1; i++ {
		curr := nums[i]
		if curr > 0 {
			break
		}
		target := 0 - curr
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]
			if target > sum {
				left++
			} else if target < sum {
				right--
			} else {
				second := nums[left]
				third := nums[right]
				v := []int{nums[i], second, third}
				ans = append(ans, v)
				//移除第一个可能重复的
				for left < right && nums[left] == second {
					left++
				}
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
