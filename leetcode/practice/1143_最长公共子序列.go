package practice

/*
*
思路:动态规划
dp[i][j]表示text1从0到i-1 与text2从0到j-1的最长公共子序列长度
dp[i][j]=dp[i-1][j-1]+1  条件:(text1[i-1]==text2[j-1])
dp[i][j]=max(dp[i-1][j],dp[i][j-1]
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
