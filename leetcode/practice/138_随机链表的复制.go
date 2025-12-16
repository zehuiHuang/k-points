package practice

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 138. 随机链表的复制

// 方法一、哈希表

/*
*
思路:创建新节点时,值很好创建,难点是Next和Random之间的关系
1、创建新节点,并且建立旧节点和新节点的MAP映射;
2、将映射的关系赋值到新节点上
*/
func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	curr := head
	for curr != nil {
		mp[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}
	curr = head

	for curr != nil {
		//例如:获取第一个新建的Node节点,那么它的下一个节点,可以通过旧节点获取新节点
		mp[curr].Next = mp[curr.Next]
		mp[curr].Random = mp[curr.Random]
		curr = curr.Next
	}
	return mp[head]
}

//方法二:
