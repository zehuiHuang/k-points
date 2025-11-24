package practice

import "math"

// d[j]表示凑齐总金额为j所需要的最少硬币数为dp[j]
// dp[j]=min(dp[j],dp[j-coin[i]]+1)

/*
*
思路：dp[j]表示为总金额j所需的最少硬币个数，那么dp[j-coins[i]]表示的是金额为j-coins[i]所需的最少硬币数,那么只需增加一个硬币：coins[i] 就是dp[j]的最少硬币数，进而推导出dp[j]=dp[j-coins[i]]+1
解释：coins[i]表示某一个硬币的面值，比如j=5，coins[i]=1(硬币面额为1),那么dp[5]一定等于dp[5-1]+1，即dp[5]=dp[4]+1
由dp[j]=dp[j-coins[i]]+1公式可知，例如j=5，那么p[5]=dp[5-coins[i]]+1，那么如何计算可以使d[5]最小呢，即遍历coins数组，找到所有可能的coins[i]，然后取最小值
即：dp[i]=min(p[5-coins[i]]+1)

由于amount也不是固定的值，所有需要遍历amount和coins数组两层遍历
最终要整理出一个二位数组：列表示的是硬币金额，行表示的是总金额
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	//为了小的不会被大的覆盖
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}
	//先物品,
	for i := 0; i < len(coins); i++ {
		//在背包
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
