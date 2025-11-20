package skill

// 思路:从0~i遍历,遍历过程中sumRight-nums[i]  sumLeft+nums[i],然后判定是否相等
func pivotIndex(nums []int) int {
	sumRight := 0
	for i := range nums {
		sumRight += nums[i]
	}
	sumLeft := 0
	for i := 0; i < len(nums); i++ {
		sumRight -= nums[i]
		if sumRight == sumLeft {
			return i
		}
		sumLeft += nums[i]
	}
	return -1
}
