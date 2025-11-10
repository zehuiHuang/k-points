package practice

// 思路:双指针
// 因为是递增的,所有可以用两个指针都放在最左边
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		v := numbers[left] + numbers[right]
		if v == target {
			return []int{left + 1, right + 1}
		} else if v > target {
			right--
		} else {
			left++
		}
	}
	return []int{0}
}
