package practice

import "math"

// 给定一个二叉树，判断它是否是 平衡二叉树
// 平衡二叉树的定义是：二叉树的每个节点的左右子树的高度差的绝对值不超过 1，则二叉树是平衡二叉树

/*
*
思路:递归,从底层开始计算左右子树的深度,若发现左右子树高度大于1,则说明不是二叉树
*/
func isBalanced(root *TreeNode) bool {
	ans := false
	var p func(root *TreeNode) int
	p = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := p(root.Left)
		right := p(root.Right)
		if math.Abs(float64(left-right)) > 1 {
			ans = true
		}
		return max(left, right) + 1
	}
	p(root)
	return ans
}
