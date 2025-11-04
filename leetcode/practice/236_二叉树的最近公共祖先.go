package practice

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/1/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/?envType=study-plan-v2&envId=top-100-liked
// 回溯算法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode

	dfs = func(root *TreeNode) *TreeNode {
		//1、终止条件(如果到叶子结点还未出现,则直接返回nil)
		if root == nil {
			return nil
		}

		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		//2、递归左右结点
		left := dfs(root.Left)
		right := dfs(root.Right)

		//3、搜集结果:直接返回
		//如果左右都存在
		if left != nil && right != nil {
			return root
		}
		//左存在,右不存在
		if left != nil && right == nil {
			return left
		}
		//左不存在,右边存在
		if right != nil && left == nil {
			return right
		}
		return nil
	}
	return dfs(root)
}
