package practice

/*
*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 回溯-深度优先搜索
func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	//方法定义
	var dfs func(root *TreeNode, targetSum int, path []int)
	dfs = func(root *TreeNode, targetSum int, path []int) {
		//判定终止条件
		if root == nil {
			return
		}
		path = append(path, root.Val)
		targetSum -= root.Val
		//什么条件下才要合并到result上?当是最后一个节点,切目标值正好减到0
		if root.Right == nil && root.Left == nil && targetSum == 0 {
			//此时为什么要用这种方式?因为切片是指针专递(复制一个副本,有相同的指针、长度和容量),防止后面再有变动而导致结果受到影响,
			//(append操作会影响切片的视野,即长度和容量),所以需要重新创建一个新的
			//举例: 如果已知有一个结果为:[1,2,3],索引地址为xxx,在后续的递归过程中,出现了[2,4,5],那么在result存储的副本就会有相同的地址、长度和容量
			// 那么此时之前被搜集的结果就会被覆盖
			newTemps := make([]int, len(path))
			copy(newTemps, path)
			result = append(result, newTemps)
		}
		//左树
		dfs(root.Right, targetSum, path)
		//右树
		dfs(root.Left, targetSum, path)
	}
	dfs(root, targetSum, []int{})
	return result
}
