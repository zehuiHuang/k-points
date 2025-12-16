package practice

/*
*
思路:使用迭代
1、先创建一个dummy节点,然后节点关系为temp -> node1 -> node2          -> node3 ->node4
2、先要转化为 temp  -> node2  -> node1                            -> node3 -> node4
3、然后将temp指针移动到node1,以此类推进行下去,直到temp为空或temp.Next 为空则停止
*/
func swapPairs(head *ListNode) *ListNode {
	//截止条件
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	temp := dummy

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		//进行交换
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		//进行移动
		temp = node1

	}
	return dummy.Next
}
