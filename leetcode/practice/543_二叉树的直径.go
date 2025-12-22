package practice

/*
*
思路:深度优先搜索
获得每个节点的深度,递归求解左右子树深度,计算左右子树深度之和(即为当前节点的直径,然后遍历过程取最大值),
返回值为左右节点的对比后的最大值,即当前节点的最大价值
*/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}
