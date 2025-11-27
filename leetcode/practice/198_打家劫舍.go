package practice

// 思路:动态规划:dp[i]=max(dp[i-2]+nums[i],dp[i-1])
func rob(nums []int) int {
	//以i为结尾房间的偷盗能获得最多金钱
	//dp[i]=max(dp[i-2]+nums[i],dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	res := 0
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
