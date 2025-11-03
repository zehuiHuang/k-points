package practice

func isSymmetric(root *TreeNode) bool {
	var f func(l *TreeNode, r *TreeNode) bool
	f = func(l *TreeNode, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil || l.Val != r.Val {
			return false
		}
		return f(l.Left, r.Right) && f(l.Right, r.Left)
	}
	return f(root.Left, root.Right)
}
