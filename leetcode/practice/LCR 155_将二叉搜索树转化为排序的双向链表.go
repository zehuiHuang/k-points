package practice

// 对应426题目
/**
思路: 关键点 有序链表、双向链表、循环链表
有序可以用中序遍历获得
双向链表:在中序遍历过程中,通过获取头节点指针head,然后移动pre,在pre和当前节点curr进行双向关联,进而形成双向链表
循环链表:将head和pre指针做关联,进行形成循环链表
*/

type NodeLink struct {
	left  *NodeLink //左子节点 对应双向链表的前节点
	right *NodeLink //右子节点 对应双向链表的下一个节点
	Val   int       //当前值
}

func treeToDoublyListNode(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	var pre *TreeNode
	var head *TreeNode

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		//下面主要是对指针pre的移动,并在移动过程中将链表关联起来,并保留head的头节点地址,作为后面循环链表的线索
		//第一次,这里一定是根节点的最下层的左子树
		if pre == nil {
			//如果pre为空,说明是第一次进来,将head指针放到头节点
			head = root
		} else {
			//如果pre不为空,则说明不是第一次进来,那么将pre(上一个节点)和当前节点关联
			pre.Right = root
		}
		//将当前节点的左节点左关联,这样就形成了双向的链表
		root.Left = pre
		//同时滑动pre到当前节点
		pre = root
		dfs(root.Right)
	}
	dfs(root)
	//最后head是在头节点,pre是在末尾节点,将首未相连几位循环链表
	head.Left = pre
	pre.Right = head
	return head
}
