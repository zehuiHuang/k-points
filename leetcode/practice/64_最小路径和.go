package practice

/**
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/
/*
*
思路:动态规划
dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
*/
func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	m := len(grid)
	n := len(grid[0])
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		//纵向最左边
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		//横向最上面
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
