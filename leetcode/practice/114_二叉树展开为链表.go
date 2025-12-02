package practice

/*
*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

/*
*
思路:
1、先进行前序遍历来收集节点
2、将节点按照要求进行处理
*/
func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		list = append(list, root)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	for i := 1; i < len(list); i++ {
		pre, current := list[i-1], list[i]
		pre.Left = nil
		pre.Right = current
	}
}
