package practice

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		lenth := len(stack)
		tmp := []int{}
		for i := 0; i < lenth; i++ {
			curr := stack[0]
			tmp = append(tmp, curr.Val)
			stack = stack[1:]
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			}
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
