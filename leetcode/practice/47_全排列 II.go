package practice

import "sort"

/*
*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],[1,2,1],[2,1,1]]

示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
//有问题待解决:todo
func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		//升序
		return nums[i] < nums[j]
	})
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
		//循环递归,每次就是对从0~n-1里选择
		//横向遍历是为了先选择其中一个
		//纵向遍历是,前面这个基础之上在选择第二个,以此类推
		for i := 0; i < n; i++ {
			//判断前一个已经相等,且已经被用过了才能不选
			if used[i] || i > 0 && nums[i] == nums[i-1] && used[i-1] {
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

func permuteUnique2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			dfs(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	dfs(0)
	return
}
