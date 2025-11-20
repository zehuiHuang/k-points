package practice

// 思路: 用二分法,中间位置为m,左右两端分别为left和right
/**
三种情况:
1、若nums[m]>nums[right],那么最小值一定落在[m+1,right]的闭区间
比如 3 4 5 6 1 2 3
nums[m]=6 大于 nums[right]=3 ,最小值1一定在m的右边
2、若nums[m]<nums[right],那么最小值一定落在[left,m]的闭区间
比如5 6 1 2 3 3 4
nums[m]=2 小于 nums[right]=4,最小值一定在m或m的左边

3、若nums[m]=nums[right],由于存在重复数,所以难判断最小值到底在左右哪个区间,
解决办法是,既然 nums[m]=nums[right],那么 那么直接right-- 来进行缩小范围即可,这样即使nums[right]是最小值,还有ums[m]在
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}
