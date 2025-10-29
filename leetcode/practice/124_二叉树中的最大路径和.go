package practice

import "math"

// 回溯算法
func maxPathSum(root *TreeNode) int {
	//定义dfs
	var dfs func(root *TreeNode) int
	res := math.MinInt32
	dfs = func(root *TreeNode) int {
		//判定条件
		if root == nil {
			return 0
		}
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)
		val := root.Val + left + right
		//收集结果(返回的是该节点的最大贡献值,就是当前节点加上左右节点中的最大值)
		res = max(res, val)
		return root.Val + max(left, right)
	}
	dfs(root)
	return res
}
