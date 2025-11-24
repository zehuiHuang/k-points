package practice

// 方法1-滚动数组
func climbStairs(n int) int {
	//a表示f(x-2) b表示f(x-1) c表示f(x)
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b     // f(x-2) = f(x-1)
		b = c     // f(x-1) = f(x)
		c = a + b // f(x) = f(x-1) + f(x-2)
	}
	return c
}

// 方法2:动态规划
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 方法3:完全背包
func climbStairs3(n int) int {
	//可以看作是背包问题
	//1和2是可选物品和对应的重量,背包为n
	//因为既可以爬1也可以爬2,可以重复,所以是完全背包问题,且要求有顺:则要求遍历是先遍历背包,在遍历物品
	//dp[j]+=dp[j-weight[i]]
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	weight := make([]int, 2)
	for i := range weight {
		weight[i] = i + 1
	}
	for j := 2; j <= n; j++ {
		for i := 0; i < 2; i++ {
			if j >= weight[i] {
				dp[j] += dp[j-weight[i]]
			}
		}
	}
	return dp[n]
}
