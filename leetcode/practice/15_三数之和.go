package practice

import "sort"

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
