package practice

/*
*
思路:24题的升级,使用迭代
1、首先创建dummy的节点,然后定义pre、start和end指针
2、先将pre、end 分别指向dummy节点
3、开始移动,start指到head节点位置,end指向start后k个节点位置
4、将start到end直接k个链表进行翻转
5、将start和end都指向pre,进行迭代处理

temp -> node1 ->node2 ->node3  -> node4 ->node5 ->node6 ,k=3
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre, end := dummy, dummy

	//dfs方法逻辑是将以head为头节点的链表进行翻转
	var dfs func(head *ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		//定义dummy节点
		var temp *ListNode
		curr := head
		//temp -> node1 ->node2 ->node3  =》  node3 -> node2 -> node1
		for curr != nil {
			next := curr.Next
			curr.Next = temp
			//移动指针
			temp = curr
			curr = next
		}
		return temp
	}

	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		//后续不够k个节点了
		if end == nil {
			break
		}
		//记录将要翻转链表的开头和结尾
		//temp -> node1 ->node2 ->node3  -> node4 ->node5 ->node6
		start := pre.Next
		next := end.Next
		//将end与后续的节点断开(为了后续start和end节点翻转做准备)
		end.Next = nil
		//将pre的下一个指向经过反转的k个节点
		pre.Next = dfs(start)

		//重新连接上,注意此时的start经过反正变成了末尾
		start.Next = next
		//对pre和end指针进行重置
		pre = start
		end = pre
	}
	return dummy.Next
}
