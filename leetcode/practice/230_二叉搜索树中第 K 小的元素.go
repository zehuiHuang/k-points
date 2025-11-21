package practice

/**
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。

输入：root = [3,1,4,null,2], k = 1
输出：1

*/

// 二叉搜索树的特定:左节点小于当前节点,当前节点小于右节点
// 思路:二叉树的中序遍历是递增的,所以进行中序遍历即可(一般的前中后序遍历特质的是根节点所在的位置)
func kthSmallest(root *TreeNode, k int) int {
	tmp := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		tmp = append(tmp, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return tmp[k-1]
}
