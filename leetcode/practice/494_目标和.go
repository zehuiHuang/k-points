package practice

/*
*
给你一个非负整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
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
