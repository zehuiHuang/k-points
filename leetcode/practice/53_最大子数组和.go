package practice

/**
leetcode:53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。


解题思路：经典动态规划，
假设结果是以i为结尾的字窜最大和为dp(i)，推导公式： dp[i]=max(dp[i - 1] + nums[i], nums[i])
1、那么首先从第2位开始，如果前面（即第1位）的值大于0，则将第一和第二位的值相加并赋值给第二位
2、从第三位开始，计算前面的（即第二位，注意：第2位的值目前是前面的累计值）是否大于0，如果大于0，则将第二位和第三位的值相加并赋值给第三位
3、以此类推，保证nums[i]的值是前面所有可能组合的最大累计值
*/

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 贪心算法
func maxSubArray2(nums []int) int {
	//以i为结尾的最大子数组合集为你dp[i]
	//dp[i]=max(dp[i - 1] + nums[i], nums[i])
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < n; i++ {
		// 这里的状态转移方程就是：求最大和
		// 会面临2种情况，一个是带前面的和，一个是不带前面的和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
	}
	return result
}
