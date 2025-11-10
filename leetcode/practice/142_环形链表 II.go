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
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
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
