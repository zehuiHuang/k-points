package practice

/*
*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]
*/

// 思路:
// 1、全排列时,需要维护一个used数组,代表已经用过了哪些数
// 2、
func permute(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	var dfs func(nums []int, path []int, used []bool)
	dfs = func(nums []int, path []int, used []bool) {
		//条件
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, path)
			return
		}
		//循环递归,就是对
		for i := 0; i < n; i++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			dfs(nums, append(path, nums[i]), used)
			//回溯时,将数据还原,path因为是复制的所有不需要还原了
			used[i] = false
		}
	}
	used := make([]bool, n)
	dfs(nums, []int{}, used)
	return ans
}
