package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路:快慢指针
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		//相等说明已经确定是环形链表
		if slow == fast {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
