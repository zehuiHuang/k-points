package dynamic_programming

import "math"

// 509. 斐波那契数
// 思路：滚动数组
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

/*
*
70. 爬楼梯：leetcode 70
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
思路：方程式：dp(i)表示爬到第i层楼梯，有d[i]个方法
理解：
1、假设知道了到达i-1阶为dp(i-1)中方法，那么对于dp(i-1)种方法中的每一种，我们都可以通过再跨1步到达第i阶，
所以说从i-1阶到i阶有dp(i-1)种方法
2、同理，假设知道了到达i-2阶为dp(i-2)中方法，那么对于dp(i-2)种方法中的每一种，我们都可以通过再跨2步到达第i阶，
所以说从i-2阶到x阶有dp(i-2)种方法
3、所以dp(i)=dp(i-1)+dp(i-2)：解释：从0爬到i-1层有dp[i-1]种方法,从0爬到i-2层有dp[i-2]种方法，从i-2层爬到i层有两种选择，i-1爬一个台阶，或者i-2爬两个台阶，其中爬一个台阶和
爬两个台阶的方式是互斥的，所以可以相加：dp(i)=dp(i-1)+dp(i-2)
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

/*
使用最小花费爬楼梯：leetcode：746
思路：推导出状态转移方程：
假定dp(i)表示爬到i阶的最小花费
1、由于是可以从0和1开始爬，那么在n=0和1的花费为0,即dp[0]=0, dp[1]=0
2、假设到达第i阶的最小花费为dp(i)，n>=2，那么dp(i) = min(dp(i-1)+cost[i-1],dp(i-2)+cost[i-2]）
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

// 62. 不同路径
func uniquePaths(m int, n int) int {
	//dp[i][j]:从0,0走到i,j一共的走法:dp[i][j]=dp[i-1][j]+dp[i][j-1]
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 416. 分割等和子集
// 思路:01背包问题
// dp[j]表示容量为j的背包最多能装多少物品
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	//背包问题
	//dp[j]=max(dp[j],dp[j-weight[i]]+value[i])
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

// 474. 一和零
// 思路:01背包,物品的重量是由两个维度来控制的
func findMaxForm(strs []string, m int, n int) int {
	//dp[i][j]:背包容量为xy(x个0和y个1)最多能装多少个物品,当前物品的重量(两个维度:x个0和y个1)
	//01背包:dp[i][j]=max(dp[i][j],dp[i-x][j-y]+1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//遍历物品
	for _, v := range strs {
		//x表示0的个数,y表示1的个数
		x, y := 0, 0
		for _, ch := range v {
			if ch == '0' {
				x++
			} else {
				y++
			}
		}
		//倒叙遍历背包
		for i := m; i >= x; i-- {
			for j := n; j >= y; j-- {
				dp[i][j] = max(dp[i][j], dp[i-x][j-y]+1)
			}
		}
	}
	return dp[m][n]
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
*
1049. 最后一块石头的重量 II
https://leetcode.cn/problems/last-stone-weight-ii/description/
思路：转化为背包问题，背包放入最大总量为total/2
*/
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

/*
*
494. 目标和 https://leetcode.cn/problems/target-sum/
*/
func findTargetSumWays(nums []int, target int) int {
	//转化为背包问题
	//left-right=target -》 right=left-target
	//由left+right=sum推到出。right=sum-left
	//left-target=sum-left -》 left=left = (target + sum)/2
	//物品为数组中的数字，背包容量为他们的值,推导 left = (target + sum)/2
	//计算数组中组合后的值为left

	//二位数组推导公式：dp[i][j]表示从0～i个范围内选取若干个物品，能够装满容量为j的背包有dp[i,j]种方法
	//推导公式：dp[i][j] = dp[i - 1][j] + dp[i - 1][j - nums[i]];

	//一维数组公式：dp[j] = dp[j] + dp[j - nums[i]]，即dp[j] += dp[j - nums[i]]
	//表示装满容量为j的背包有dp[j]种方法

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bagSize := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bagSize+1)

	// 初始化，不放物品也算是一种方法
	dp[0] = 1
	//遍历：先物品，在容量，且容量倒序遍历
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

/*
硬币兑换2 leetcode：518
https://leetcode.cn/problems/coin-change-ii/description/
思路：

*
*/
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
		if j%coins[0] == 0 {
			dp[0][j] = 1
		}
	}
	//最左边的物品列初始化
	for i := 0; i < len(coins); i++ {
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

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	//定义以i为结构的最长子系列长度为dp[i]
	//dp[i]=max(dp[i],dp[j]+1) v[i]>v[j]的值
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			//计算以i为结尾的最长递增子序列，只要满足：v[i]>v[j]，那么就是计算出0～j之间每个子序列长度+1就是以
			//i为结尾的最长子序列
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

// 最长递增子序列
func lengthOfLIS2(nums []int) int {
	//dp[i]:以i为结尾的最长递增子序列的长度为dp[i]
	//递推公式:dp[i]=
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// 213. 打家劫舍 II
func rob(nums []int) int {
	//dp[i]=max(dp[i-2]+nums[i],dp[i-1])
	var ff func(nums []int) int
	ff = func(nums []int) int {
		dp := make([]int, len(nums))
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < len(nums); i++ {
			dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		}
		return dp[len(nums)-1]
	}
	if len(nums) == 1 {
		return nums[0]
	}
	d1 := ff(nums[1:])
	d2 := ff(nums[:len(nums)-1])
	return max(d1, d2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 337. 打家劫舍 III
func rob3(root *TreeNode) int {
	//思路,从底向上遍历(后序遍历:即递归),定义dfs,返回值分别表示选当前节点和不选当前节点的
	var dfs func(root *TreeNode) (a, b int)

	dfs = func(root *TreeNode) (a, b int) {
		if root == nil {
			return 0, 0
		}
		//左儿子选or不选:l1是选的值,l2是不选的值
		l1, l2 := dfs(root.Left)
		//右儿子选or不选:r1是选的值,r2是不选的值
		r1, r2 := dfs(root.Right)
		//选择当前节点:表示不选左儿子的值+不选右儿子的值
		v1 := root.Val + l2 + r2
		//不选当前节点
		v2 := max(l1, l2) + max(r1, r2)
		return v1, v2
	}
	return max(dfs(root))
}

// 718. 最长重复子数组
func findLength(nums1 []int, nums2 []int) int {
	ans := 0
	//dp[i][j]表示以i-1结尾的nums1和以j-1结尾的nums2的最长重复字数组的长度为dp[i][j]
	dp := make([][]int, len(nums1)+1)
	for i := 0; i <= len(nums1); i++ {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans
}
func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j]表示text1从0~i-1 余text2从0~j-1的最长公共子序列长度
	//dp[i][j]=dp[i-1][j-1]+1  条件:(text1[i-1]==text2[j-1])
	//dp[i][j]=max(dp[i-1][j],dp[i][j-1]
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
