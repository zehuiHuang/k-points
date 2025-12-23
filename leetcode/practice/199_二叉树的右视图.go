package practice

/*
*
思路:层次遍历,等层次遍历时,每一层的最后一位便是能看到到
*/
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	num := 1
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			num--
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if num == 0 {
				ans = append(ans, node.Val)
				num = len(queue)
			}
		}
	}
	return ans
}
