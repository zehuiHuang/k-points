package linked

import (
	"container/list"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
单向链表反转
*/
func revertList(header *ListNode) *ListNode {
	if header.Next == nil {
		return header
	}
	var next *ListNode
	var pre *ListNode
	for header != nil {
		//记录下一个节点的地址
		next = header.Next
		//把当前节点的指针指向前一个节点
		header.Next = pre
		//前一个节点指向当前节点
		pre = header
		header = next
	}
	return pre
}

type Node struct {
	Val   int
	left  *Node
	right *Node
}

/*
*
单向链表前序遍历
*/
func preOrder(header *Node) {
	if header == nil {
		return
	}
	print(header.Val)
	preOrder(header.left)
	preOrder(header.right)
}

func preorderTraversal(root *Node) (vals []int) {
	var preorder func(*Node)
	preorder = func(node *Node) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.left)
		preorder(node.right)
	}
	preorder(root)
	return
}

/*
*
单向链表中序遍历
*/
func inOrder(header *Node) {
	if header == nil {
		return
	}
	preOrder(header.left)
	print(header.Val)
	preOrder(header.right)
}

/*
*
单向链表后序遍历
*/
func blackOrder(header *Node) {
	if header == nil {
		return
	}
	preOrder(header.left)
	preOrder(header.right)
	print(header.Val)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathTarget(root *Node, sum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	path := []int{}
	ans = process(root, path, 0, sum)
	return ans
}

func process(x *Node, path []int, pathSum int, target int) (ans [][]int) {
	//为叶子结点，则判定路径加上当前结果val是否等于目标值
	if x.left == nil && x.right == nil {
		if pathSum+x.Val == target {
			path = append(path, x.Val)
			ans = append(ans, path)
		}
		return ans
	}
	//非叶子结点，继续递归
	path = append(path, x.Val)
	if x.left != nil {
		ans = process(x.left, path, pathSum+x.Val, target)
	}
	if x.right != nil {
		ans = process(x.right, path, pathSum+x.Val, target)
	}
	path = path[:len(path)-1]
	return ans
}

// leedcode 2 https://leetcode-cn.com/problems/add-two-numbers/
/**
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	//进位
	carry := 0
	for l1 != nil || l2 != nil {
		//当前两个链表结点的值
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//当前两个结点的值加上进位
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			//新建一个结点
			tail.Next = &ListNode{Val: sum}
			//指向下一个结点
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

/*
*
* 合并两个有序链表 leetcode 21
思路：
1、如果两个链表都为空，则返回空
2、如果其中一个链表为空，则返回另一个链表
3、如果两个链表都不为空，则比较两个链表的头结点的值，较小的那个结点指向下一个结点，然后递归调用合并函数
4、返回合并后的链表
*/
// mergeTwoLists 合并两个有序链表，返回合并后的有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 如果list1为空，直接返回list2（递归终止条件）
	if list1 == nil {
		return list2
		// 如果list2为空，直接返回list1（递归终止条件）
	} else if list2 == nil {
		return list1
		// 如果list1的值小于list2的值，将list1的下一个节点与list2合并，结果作为list1的下一个节点
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
		// 否则，list2的值小于等于list1的值，将list2的下一个节点与list1合并，结果作为list2的下一个节点
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	prehead := &ListNode{}
	prev := prehead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	// 合并后 l1 和 l2 最多只有一个还未被合并完，我们直接将链表末尾指向未合并完的链表即可
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}
	return prehead.Next
}

/*
*
leetcode 105 从前序与中序遍历序列构造二叉树
思路，递归，通过前序遍历找到根节点，然后在中序遍历中找到根节点的位置，左边的是左子树，右边的是右子树
*/
func buildTree(preorder []int, inorder []int) *Node {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}
	preorderMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		preorderMap[inorder[i]] = i
	}
	return buildTree2(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, preorderMap)
}

func buildTree2(preorder []int, l1, r1 int, inorder []int, l2, r2 int, preorderMap map[int]int) *Node {
	if l1 > r1 {
		return nil
	}
	if l1 == r1 {
		return &Node{
			Val: preorder[l1],
		}
	}
	//根节点
	root := &Node{
		Val: preorder[l1],
	}
	/**
	1 2. 4 5 3 6 7
	4 2 5 1 6 3 7
	*/
	//找到中序遍历的根节点位置（下标位置）
	/**
	  find=3
	  1,3
	*/
	find := preorderMap[preorder[l1]] //3
	//长度为find-l2
	root.left = buildTree2(preorder, l1+1, l1+find-l2, inorder, l2, find-1, preorderMap)
	root.right = buildTree2(preorder, l1+find-l2+1, r1, inorder, find+1, r2, preorderMap)
	return root
}

/*
*
160. 相交链表
描述：给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
思路：使用双指针，并将两个链表连起来，遍历链表时第一个相同的就是相交节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果任一链表为空，则不可能有相交节点
	if headA == nil || headB == nil {
		return nil
	}

	// 初始化两个指针，分别指向两个链表的头节点
	pa, pb := headA, headB

	// 使用双指针法遍历两个链表
	// 当两个指针相遇时，即为相交节点
	// 如果两个链表不相交，最终两个指针都会指向nil
	for pa != pb {
		// 如果pa到达链表A的末尾，则从链表B的头部继续遍历
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		// 如果pb到达链表B的末尾，则从链表A的头部继续遍历
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	// 返回相交节点（如果存在）或nil（如果不存在）
	return pa
}

/*
*
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
思路：定义pre、curr指针，pre指向空节点，curr指向头节点，头节点指向pre时，
遍历时，为了保证Next指针不丢失， 需要先并保存curr的Next指针，然后curr的Next指到pre，最后右移动pre和curr
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	//当前节点为头节点
	curr := head
	//当前节点不为空则继续循环
	for curr != nil {
		//保存当前节点的Next节点
		next := curr.Next
		//当前节点的下一个节点赋值为pre
		curr.Next = prev
		//将pre指针移动到curr节点
		prev = curr
		//将curr指针移动到当前节点的Next节点
		curr = next
	}
	return prev
}

/*
*
234. 回文链表
*/
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	//转成数组
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	ans := []int{}
	for ; head != nil; head = head.Next {
		ans = append(ans, head.Val)
	}
	l, r := 0, len(ans)-1
	for l < r {
		if ans[l] != ans[r] {
			return false
		}
		l++
		r--
	}
	return true
}

/*
*146. LRU 缓存
 */
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/*
*
迭代法前序遍历
*/
func preorderTraversal2(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*Node)

		ans = append(ans, node.Val)
		if node.right != nil {
			st.PushBack(node.right)
		}
		if node.left != nil {
			st.PushBack(node.left)
		}
	}
	return ans
}

func levelOrder(root *Node) [][]int {
	st := list.New()
	st.PushBack(root)
	ans := [][]int{}
	for st.Len() > 0 {
		length := st.Len()
		arr := []int{}
		for i := 0; i < length; i++ {
			node := st.Remove(st.Back()).(*Node)
			arr = append(arr, node.Val)
			if node.left != nil {
				st.PushBack(node.left)
			}
			if node.right != nil {
				st.PushBack(node.right)
			}
		}
		ans = append(ans, arr)
	}
	return ans
}

func levelOrder2(root *Node) [][]int {
	res := [][]int{}
	if root == nil { //防止为空
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len() //保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node) //出队列
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
			tmpArr = append(tmpArr, node.Val) //将值加入本层切片中
		}
		res = append(res, tmpArr) //放入结果集
		tmpArr = []int{}          //清空层的数据
	}

	return res
}

// 226. 翻转二叉树
// 思路：利用前序遍历
func invertTree(root *Node) *Node {
	var p func(root *Node)
	p = func(root *Node) {
		if root == nil {
			return
		}
		left := root.left
		root.left = root.right
		root.right = left
		p(root.left)
		p(root.right)
	}
	p(root)
	return root
}

// 236. 二叉树的最近公共祖先
// 思路：使用回溯
func lowestCommonAncestor(root, p, q *Node) *Node {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)
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

func isBalanced(root *Node) bool {
	ans := false
	var p func(root *Node) int
	p = func(root *Node) int {
		if root == nil {
			return 0
		}
		left := p(root.left)
		right := p(root.right)
		if math.Abs(float64(left-right)) > 1 {
			ans = true
		}
		return max(left, right) + 1
	}
	p(root)
	return ans
}
