package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
1
23
45 66
*/
// 103. 二叉树的锯齿形层序遍历
//思路,在层序遍历的基础上,对每层的tmp数据根据奇偶进行翻转即可
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	level := 0
	for len(stack) > 0 {
		length := len(stack)
		tmp := []int{}
		level++
		for i := 0; i < length; i++ {
			v := stack[0]
			stack = stack[1:]

			tmp = append(tmp, v.Val)
			if v.Left != nil {
				stack = append(stack, v.Left)
			}
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
		}
		// 修复：应该反转当前层的输出结果 tmp，而不是下一层的输入 stack
		// 奇数层（从1开始计数）需要反转输出顺序
		if level%2 == 0 {
			l, r := 0, len(tmp)-1
			for l < r {
				tmp[l], tmp[r] = tmp[r], tmp[l]
				l++
				r--
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}
