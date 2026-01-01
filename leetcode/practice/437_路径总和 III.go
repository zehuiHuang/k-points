package practice

/*
*
思路:
1、利用前缀和+map
2、利用回溯遍历,搜集符合条件的节点
3、遍历完子节点后,要恢复现场(因为累计和时要求必须是 一个节点是另外一个节点的祖先)
*/
func pathSum3(root *TreeNode, targetSum int) int {
	ans := 0
	//k为从根节点到当前节点的累计和,value为累计和出现的次数
	mp := make(map[int]int)
	//可以理解为:
	mp[0] = 1
	//node 为当前节点,s为从根节点到当前节点的父节点的累积和
	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		//终止条件
		if node == nil {
			return
		}
		//从根节点到当前节点的累积和
		s += node.Val

		ans += mp[s-targetSum]

		mp[s] += 1
		dfs(node.Left, s)
		dfs(node.Right, s)
		//恢复现场
		mp[s] -= 1
	}

	dfs(root, 0)
	return ans
}
