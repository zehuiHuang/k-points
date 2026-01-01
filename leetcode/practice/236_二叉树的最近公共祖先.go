package practice

// 回溯算法
/**
思路:
1、回溯,从最底层开始查找
假设在最底层时:
如果该节点等于目标节点p或者q,则直接朝上返回其本身
如果都不等,则直接返回空即可

2、那么在倒数第二层时,判断左右节点是否为空的情况,如果都不为空,说明当前节点就是最近的公共祖先,并朝上返回当前节点即可,并一路返回到顶
如果左不为空右为空,则将左直接返回即可, 左为空右不为空,则返回右节点即可

3、需要注意的是,如果存在p的子节点就是q,那么在代码下探的时候,在遇到p是直接就返回了(直接到顶),q也没机会进行左右递归,这样p就是最近的公共祖先,也符合题目
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//入参为根节点root,返回参数为p或q的最新公共祖先
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
		left := dfs(root.Left)   //返回3
		right := dfs(root.Right) //返回空

		//3、搜集结果:直接返回
		//如果左右都存在(当前节点便是公共祖先)
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

// 思路：使用递归
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
