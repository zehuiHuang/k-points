package practice

/*
*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
*/

// 思路:
// 1、通过前序遍历找到根节点,然后通过中序遍历,找到根节点所在的index,index左边都是左子树的节点,右边都是右子树的节点
// 2、由此可以得到根节点、根节点的左节点,根节点的右节点
// 3、最后在子树上重复以上操作即可

// todo
func buildTree(preorder []int, inorder []int) *TreeNode {
	mp := make(map[int]int)
	//中序遍历的节点值所在索引;kv value:index
	for i := 0; i < len(inorder); i++ {
		mp[inorder[i]] = i
	}
	//root:前序遍历列表的头节点
	//left:中序遍历列表中的左(或右)子树位置起始位置
	//right:中序遍历列表中的左(或右)子树位置的末尾
	var dfs func(root, left, right int) *TreeNode

	dfs = func(root, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		node := &TreeNode{
			Val: preorder[root],
		}
		//中序遍历的index左边为左子树,右边为右子树,两边可以进行左右子树划分
		index := mp[preorder[root]]
		//开始左子树递归
		//第一个参数为头节点下一位,也就是左子树的头节点
		//第二个参数  left 是指中序遍历列表中的左边
		//第三个参数  i-1 是中序遍历列表中左子树的最右边
		//整个左子树在中序遍历列表中的范围为:left~i-1
		node.Left = dfs(root+1, left, index-1)
		//求解前序遍历中右子树的根节点位置:
		/**
		1、先计算左子树长度
		在中序数组中，左子树的范围是 [left, index-1]
		左子树节点个数 = (index-1) - left + 1 = index - left

		2、计算前序数组中右子树根节点位置
		前序遍历顺序：根 → 左子树 → 右子树
		当前根节点位置：root
		左子树占据的位置：从 root+1 到 root + (index-left)
		因此右子树根节点位置 = root + (index-left) + 1
		*/
		node.Right = dfs(root+(index-left)+1, index+1, right)
		return node
	}
	return dfs(0, 0, len(inorder)-1)
}
