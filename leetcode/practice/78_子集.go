package practice

/*
*
思路:使用回溯算法

1、按照步骤进行回溯算法定义
2、i定义为nums[i]当前的值选或者不选
3、终止条件为i==n
4、回溯时,一种为不选nums[i],则直接dfs(i+1, path),另外一种为选,则dfs(i+1, append(path, nums[i]))
*/
func subsets(nums []int) [][]int {
	n := len(nums)
	//我们知道,一个数选或者不选有两种情况,那么一共的子集的数量一定是2的n次方
	ans := make([][]int, 0, 1<<n)

	var dfs func(int, []int)
	//i 表示nums[i]当前的值选或者不选
	dfs = func(i int, path []int) {
		//终止条件
		if i == n {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		//不选 nums[i]
		dfs(i+1, path)
		//选 nums[i]
		dfs(i+1, append(path, nums[i]))
	}
	dfs(0, []int{})
	return ans
}
