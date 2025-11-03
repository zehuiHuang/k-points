package practice

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
