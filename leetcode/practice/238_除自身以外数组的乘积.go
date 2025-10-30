package practice

// 技巧类型的题目
func productExceptSelf(nums []int) []int {
	n := len(nums)
	//思路:从左边开始先计算i左边所有数的乘积,从右边开始,计算i右边所有数的乘积, 顺便统计answer[i]=左边+右边
	answer := make([]int, len(nums))
	//起始位置左边没有数,左右=1
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	//从右边开始计算
	//从右边开始,他的右边没值,所有=1
	R := 1
	for i := n - 1; i >= 0; i-- {
		//answer[i]=左乘积+右乘积
		answer[i] = answer[i] * R
		//计算下一个位置的右边的所有乘积
		R = nums[i] * R
	}
	return answer
}
