package practice

/*
使用最小花费爬楼梯：leetcode：746
思路：推导出状态转移方程：
假定dp(i)表示爬到i阶的最小花费
1、由于是可以从0和1开始爬，那么在n=0和1的花费为0,即dp[0]=0, dp[1]=0
1、假设到达第i阶的最小花费为dp(i)，n>=2，那么dp(i) = min(dp(i-1)+cost[i-1],dp(i-2)+cost[i-2]）
*/

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
