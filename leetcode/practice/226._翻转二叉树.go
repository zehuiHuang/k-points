package practice

func invertTree(root *TreeNode) *TreeNode {
	var f func(root *TreeNode)

	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		//前序遍历
		node := root.Left
		root.Left = root.Right
		root.Right = node
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return root
}
