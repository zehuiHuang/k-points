package practice

/*
*
思路:递归,
1、定义函数入参:左节点和右节点:判定左右节点是否对称(值是否相等)
2、然后在让各自左右节点下的子左右节点进行对比,对比规则为:
左子节点的左节点和右子节点的右节点进行对称判定
右子节点的左节点和左子节点的右节点进行对称判定

3、判定对称规则为:值相等或都为空
*/
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
