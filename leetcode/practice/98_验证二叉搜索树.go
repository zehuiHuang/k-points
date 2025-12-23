package practice

import "math"

/*
*
思路:利用中序遍历的顺序性,判断当前节点的值是否比前一个节点的值小
*/
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) { // 左
			return false
		}
		if node.Val <= pre { // 中
			return false
		}
		// 更新前一个节点,为后面与下一个节点进行数值比较
		pre = node.Val
		return dfs(node.Right) // 右
	}
	return dfs(root)
}
