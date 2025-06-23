package dynamic_programming

import "math"

/**
leetcode:53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。


解题思路：经典动态规划，
假设结果是以i为结尾的字窜最大和为f(i)，
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

/*
*
70. 爬楼梯：leetcode 70
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

思路：方程式：dp(i)表示爬到第i层楼底，有d[i]个方法
理解：
1、假设知道了到达i-1阶为dp(i-1)中方法，那么对于dp(i-1)种方法中的每一种，我们都可以通过再跨1步到达第i阶，
所以说从i-1阶到i阶有dp(i-1)种方法
2、同理，假设知道了到达i-2阶为f(i-2)中方法，那么对于dp(i-2)种方法中的每一种，我们都可以通过再跨2步到达第i阶，
所以说从i-2阶到x阶有dp(i-2)种方法
3、所以dp(i)=f(i-1)+f(i-2)：解释：从0爬到i-1层有dp[i-1]种方法,从0爬到i-2层有dp[i-2]种方法，从i-2层爬到i层有两种选择，i-1爬一个台阶，或者i-2爬两个台阶，其中爬一个台阶和
爬两个台阶的方式是互斥的，所以可以相加：dp(i)=f(i-1)+f(i-2)
*/
func climbStairs(n int) int {
	// 初始化三个变量：
	// p: 表示dp(i-2)，即前两步的方法数
	// q: 表示dp(i-1)，即前一步的方法数
	// r: 表示dp(i)，即当前步的方法数
	p, q, r := 0, 0, 1
	// 从第1阶开始计算到第n阶
	for i := 1; i <= n; i++ {
		p = q     // dp(i-2) = dp(i-1)
		q = r     // dp(i-1) = dp(i)
		r = p + q // dp(i) = dp(i-1) + dp(i-2)
	}
	return r
}

/*
斐波那契数：leetcode 509
//思路：滚动数组
*/
func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}

/*
使用最小花费爬楼梯：leetcode：746
思路：推导出状态转移方程：
假定f(i)表示爬到i阶的最小花费
1、由于是可以从0和1开始爬，那么在n=0和1的花费为0,即dp[0]=0, dp[1]=0
1、假设到达第i阶的最小花费为f(i)，n>=2，那么f(i) = min(f(i-1)+cost[i-1],f(i-2)+cost[i-2]）
*/

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(pre+cost[i-2], cur+cost[i-1])
	}
	return cur
}

/*
01背包问题
dp[i][j]表示 从0～i的物品任意取，放进容量为j的背包，计算它的最大价值，dp[i][j]即表示他的最大价值
**/

/*
零钱兑换 leetcode 322
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

思路：dp[j]表示为总金额j所需的最少硬币个数，那么dp[j-coins[i]]表示的是金额为j-coins[i]所需的最少硬币数,那么只需增加一个硬币：coins[i] 就是dp[j]的最少硬币数，进而推导出dp[j]=dp[j-coins[i]]+1
解释：coins[i]表示某一个硬币的面值，比如j=5，coins[i]=1(硬币面额为1),那么dp[5]一定等于dp[5-1]+1，即dp[5]=dp[4]+1
由dp[j]=dp[j-coins[i]]+1公式可知，例如j=5，那么p[5]=dp[5-coins[i]]+1，那么如何计算可以使d[5]最小呢，即遍历coins数组，找到所有可能的coins[i]，然后取最小值
即：dp[i]=min(p[5-coins[i]]+1)

由于amount也不是固定的值，所有需要遍历amount和coins数组两层遍历
最终要整理出一个二位数组：列表示的是硬币金额，行表示的是总金额
**/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//总金额为0，那么硬币个数也为0
	dp[0] = 0
	//初始化：防止dp[j]=min(dp[j-coins[i]]+1,dp[j]) 比较时被覆盖
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	n := len(coins)
	//先遍历硬币，在遍历金额(遍历物品)
	for i := 0; i < n; i++ {
		//遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// coins = [1, 2, 5], amount = 11
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//总金额为0，那么硬币个数也为0
	dp[0] = 0
	//初始化dp[i]=math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}
	//[0,math.MaxInt32,math.MaxInt32,math.MaxInt32,math.MaxInt32]

	//先遍历硬币，在遍历金额
	//遍历硬币
	for i := 0; i < len(coins); i++ {
		//遍历金额
		for j := coins[i]; j <= amount; j++ {
			//公式：dp[j]: dp[j-coins[i]]+1
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

/*
硬币兑换2 leetcode：518
https://leetcode.cn/problems/coin-change-ii/description/
思路

**/

func change(amount int, coins []int) int {
	return 0
}

func lastStoneWeight(stones []int) int {
	total := 0
	for i := 0; i < len(stones); i++ {
		total += stones[i]
	}
	//转化为背包问题：分成两份尽量一样的石头，则total/2表示背包容量，石头表示物品
	//推到公式：dp[j]=max(dp[j],dp[j-stones[i]]+stones[i])，表示容量为j的背包所用放入的最大（价值）重量

	//初始化，默认为零

	target := total / 2
	dp := make([]int, total)
	//遍历：先遍历物品，在倒序遍历背包
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return total - 2*dp[target]
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
