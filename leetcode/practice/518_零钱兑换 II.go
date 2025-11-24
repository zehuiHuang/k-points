package practice

/*
*
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。

请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。

假设每一种面额的硬币有无限个。

题目数据保证结果符合 32 位带符号整数。

示例 1：

输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
*/

// 思路:组合数的dp结构为:dp[i][j]表示从0~i选择物品,能将容量为j的背包装满的组合数为dp[i][j]
// 因为硬币数量无限个,所有是无限背包
// dp[i][j]=dp[i-1][j]+dp[i][j-nums[i]]

func change(amount int, coins []int) int {
	//amount是背包容量，金币是物品
	//推算公式：dp[i][j]=dp[i-1][j]+dp[i][j-coins[i]] 表示从下标0～i之间选择硬币，凑满j的总金额的组合数为dp[i][j]
	//那么组合数等于 放 i 硬币的组合数+不放i硬币的组合数
	//转一纬数组公式:dp[j] +=dp[j-coins[i]]

	//初始化
	dp := make([][]int, len(coins))

	for j := 0; j < len(dp); j++ {
		dp[j] = make([]int, amount+1)
	}
	dp[0][0] = 0
	//最上面的一列 容量行初始化
	for j := 1; j <= amount; j++ {
		//能被整除则表示有一种方案
		if j%coins[0] == 0 {
			dp[0][j] = 1
		}
	}
	//最左边的物品列初始化
	for i := 0; i < len(coins); i++ {
		//用物品i装满容量为0的背包,就是不放物品,那么都算一种
		dp[i][0] = 1
	}
	//遍历，先遍历物品，在遍历容量
	for i := 1; i < len(coins); i++ {
		//遍历容量
		for j := 0; j <= amount; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(coins)-1][amount]
}
