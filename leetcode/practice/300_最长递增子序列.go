package practice

func lengthOfLIS(nums []int) int {
	//定义以i为结尾的最长子系列长度为dp[i]
	//dp[i]=max(dp[i],dp[j]+1)  条件:v[i]>v[j]的值
	dp := make([]int, len(nums))
	for i := range dp {
		//因为最短也就是1
		dp[i] = 1
	}
	ans := dp[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			//0.....j...i
			//计算以i为结尾的最长递增子序列，只要满足：v[i]>v[j]，那么就是计算出0～j之间每个子序列长度+1就是以i为结尾的最长子序列
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}
