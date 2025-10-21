package practice

// 回溯算法
func predictTheWinner(nums []int) bool {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	//1、定义递归函数:start和end分别表示数字的范围,start从下表为零开始,首选的人获得的总积分
	var dfs func(nums []int, start, end int) int
	dfs = func(nums []int, start, end int) int {
		//2、结束条件
		if end == start {
			return nums[end]
		}
		//两种方案:nums长度大于2 或者小于等于2
		var left, right int
		//0 1 2
		if end-start >= 2 {
			//dfs(nums, start+2, end)表示对手选择了左边,可以被先手选择的区间变为从start+2-> end
			//dfs(nums, start+1, end-1)表示对手选择了右边,可以被先手选择的区间变为start+1->end—1
			//取最小值说明对手也很聪明,也选择可以获得积分最多的选择方式
			left = nums[start] + min(dfs(nums, start+2, end), dfs(nums, start+1, end-1))
			//和left原理相同
			right = nums[end] + min(dfs(nums, start, end-2), dfs(nums, start+1, end-1))
			return max(left, right)
		} else {
			return max(nums[start], nums[end])
		}
	}
	firstSelected := dfs(nums, 0, len(nums)-1)
	secondSelected := total - firstSelected
	return firstSelected >= secondSelected
}

// 贪心算法
func predictTheWinner2(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	//1、定义状态转移方程
	//dp[i][j]表示在i~j这个数据区间,先手能赢对手的差值
	//dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
	//解释:先手如果选择左边,则对手可选范围为i+1->j,则对手能赢先手的差值便为dp[i+1][j],那么nums[i]-dp[i+1][j]就是先手能赢对手的差值
	//如果先手选择右边,则对手可选范围为i->j-1,则对手能赢先手的差值便为dp[i][j-1],那么nums[j]-dp[i][j-1]就是先手能赢对手的差值
	//最后这俩选择最大值,即选择先手能赢对手的差值的最大值(贪心)
	//2、初始化
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}
	//3、遍历
	//例子:[1,5,233,7]
	//解释:
	/**
	第一层循环i值:
	i:=2~0-> 2, 1, 0
	第二次循环j值:
	j:=3,2,1
	方程转移值:
	dp[2][3]
	dp[1][2]、dp[1][3]
	dp[0][1]、dp[0][1]、dp[0][3]
	*/
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}
