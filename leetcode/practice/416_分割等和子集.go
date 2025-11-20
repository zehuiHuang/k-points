package practice

// 使用01背包二维数组推导公式
// 思路:01背包问题
// dp[i][j]表示从0~i的物品里任意取,放到容量为j的背包里,价值的总和为dp[i][j]
// dp[i][j]=max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]int, len(nums))
	//物品
	for i := range dp {
		//初始化列数时,包容量+1是因为 背包大小为target时,有背包背包为0,所有他的列数应该是target+1个
		dp[i] = make([]int, target+1)
	}
	//初始化:默认都是0

	//先物品,i必须从1开始,因为物品为0没啥意义,不需要遍历
	for i := 1; i < len(nums); i++ {
		//背包
		for j := 0; j <= target; j++ {
			//存在一种情况:物品的大小背包,则背包就不能放入了
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}
	//
	return dp[len(nums)-1][target] == target
}

// 使用01背包一维数组推导公式：
func canPartition2(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	//dp[j]=max(dp[j],dp[j-weight[i]]+value[i])  表示将容量j的背包最多能装多少物品
	dp := make([]int, target+1)
	//一维数组需要注意,一定要先遍历物品,在遍历背包,切背包要倒叙
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
