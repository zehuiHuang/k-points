package practice

/*
*
思路:分治思想

对sortList方法进行递归调用:
递归过程包括:1)将链表从中中间分开,获得两个子链表;2)对两个链表分别进行递归调用,递归的sortList返回子链表的有序链表的合并.

递归过程中,最大的子问题为连个链表长度为1,并对此进行有序链表合并,保证sortList返回的是有序链表,
在后续的递归中,都是将有序链表进行合并的过程.
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//1、找到中间节点并将链表分成两部分
	mid := findMidAndSplit(head)
	//2、将分割的两个链表分别进行递归
	left := sortList(head) // 前半部分
	right := sortList(mid) // 后半部分
	//3、合并有序链表
	return mergeListNode(left, right)
}

// 利用快慢指针找到中间节点,并对链表进行拆分
func findMidAndSplit(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	prev := head

	// 快慢指针找到中点
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开链表
	prev.Next = nil
	return slow
}

// 合并有序队列
func mergeListNode(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	dummy := &ListNode{0, nil}
	curr := dummy

	// 循环比较合并
	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	// 连接剩余节点
	if left != nil {
		curr.Next = left
	} else {
		curr.Next = right
	}

	return dummy.Next
}
