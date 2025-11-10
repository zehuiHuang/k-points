package practice

// 层序遍历
func maxDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		ans++
		for i := 0; i < n; i++ {
			node := queue[0]
			//截取
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}

// 递归
// 思路,定义dfs,返回该节点的贡献(层数)
func maxDepth2(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		return max(left, right) + 1
	}
	return dfs(root)
}
