package practice

/*
*
动态规划:dp[i]=max(dp[i-2]+nums[i],dp[i-1]),以i为结尾房间的偷盗能获得最多金钱时dp[i]
另外因为是一个环,所有在选择时有两种情况:
第一种:偷第一家,那么最后一家就不能偷盗
第二种:偷最后一家,那么第一家就不能偷盗
把这俩情况都做了,然后比对最大值即可
*/
func rob22(nums []int) int {
	//dp[i]=max(dp[i-2]+nums[i],dp[i-1])
	var ff func(nums []int) int
	ff = func(nums []int) int {
		dp := make([]int, len(nums))
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
		return dp[len(nums)-1]
	}
	if len(nums) == 1 {
		return nums[0]
	}
	d1 := ff(nums[1:])
	d2 := ff(nums[:len(nums)-1])
	return max(d1, d2)
}
