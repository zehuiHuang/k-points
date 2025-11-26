package practice

import "math"

// 给定一个二叉树，判断它是否是 平衡二叉树
// 平衡二叉树的定义是：二叉树的每个节点的左右子树的高度差的绝对值不超过 1，则二叉树是平衡二叉树

/*
*
思路:递归,从底层开始计算左右子树的深度,若发现左右子树高度大于1,则说明不是二叉树
*/
func isBalanced(root *TreeNode) bool {
	ans := true
	var f func(node *TreeNode) int
	f = func(node *TreeNode) int {
		if node == nil || !ans {
			return 0
		}
		l := f(node.Left)
		r := f(node.Right)
		if math.Abs(float64(l-r)) > 1 {
			ans = false
			return 0
		}
		return max(l, r) + 1
	}
	f(root)
	return ans
}
