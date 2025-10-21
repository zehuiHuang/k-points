package practice

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
			//此时为什么要用这种方式?因为切片是指针专递,防止后面再有变动而导致结果受到影响,所以需要重新创建一个新的
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
