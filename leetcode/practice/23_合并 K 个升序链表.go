package practice

/*
*
思路:分治合并最清晰
*/
func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for _, list := range lists {
		ans = mergeListNode(ans, list)
	}
	return ans
}
