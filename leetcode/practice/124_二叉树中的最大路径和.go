package practice

import "math"

/**
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和
*/

// 回溯算法

/*
*
思路:在回溯过程中,
1、获得当前节点的最大贡献值,即当前节点加上左右节点中的最大值,返回该节点的最大贡献值
2、每次以当前节点作为根节点的最大路径和(val+left.val+right.val),跟换结果(res)进行比较,取最大值
、
*/
func maxPathSum(root *TreeNode) int {
	//定义dfs:参数为根root,返回其 最大路径和
	var dfs func(root *TreeNode) int
	res := math.MinInt32
	dfs = func(root *TreeNode) int {
		//判定条件
		if root == nil {
			return 0
		}
		//如果左右节点返回的是负数,那么可以排出掉,因为负数加上当前节点的值,会变小
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)
		//以当前节点作为根节点的最大路径和
		val := root.Val + left + right
		//跟其他以自身节点为根节点的最大路径和做比较,取最大值
		res = max(res, val)
		//收集结果(返回的是该节点的最大贡献值,就是当前节点加上左右节点中的最大值)
		return root.Val + max(left, right)
	}
	dfs(root)
	return res
}
