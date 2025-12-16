package practice

/*
*
思路:主要是找到倒数第n个节点
三种:
第一种:求出总长度m,然后m-n的位置即为其下一个节点是要删除的
第二种: 配合先进后出的栈来找到倒数第n个节点
第三种:双指针,让双指针间隔n,当前面的指针为nil了说明后面的指针对应的节点就是要删除的节点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//例如:dummy ->  1 -> 3 -> 6 -> 4 -> 9 ; n=2

	//本题选择第三种方式
	dummy := &ListNode{0, head}
	first, second := head, dummy
	//现将一个指针移动到n位置
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	//删除second.Next节点,也就是倒数第n个节点
	second.Next = second.Next.Next
	return dummy.Next
}
