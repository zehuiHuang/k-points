package practice

/**
给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。

返回 你可以获得的最大乘积 。

示例 1:

输入: n = 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: n = 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
*/

/*
*
思路:动态规划
1、定义dp,dp[i]表示将正整数i拆分成至少两个正整数的和时,这些正整数乘积最大值为dp[i]

2、状态转移:
将i拆分成j和i-j,然后i-j可以继续拆分
1)如果i-j的值不继续拆分,则dp[i]=j*(i-j)
2)如果i-j的值继续拆分,则dp[i]= j*dp[i-j]
此时 只需取两种的最大值即可: dp[i]=max(j*(i-j),j*dp[i-j])
只要j固定了,那么上式子就永远成立
因为j的取值范围:1<=j<i,所以遍历将所有j的情况统计出来,再取最大值即可
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		//以i为结尾的所有情况
		curMax := 0
		for j := 1; j < i; j++ {
			//以i为结尾的 从1到i的所有j的情况
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
