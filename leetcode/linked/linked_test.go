package linked

import "testing"

func TestName(t *testing.T) {
	ListNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	ListNode = revertList(ListNode)
	print(ListNode)
}

// 1245367
func TestPreOrder(t *testing.T) {
	header := &Node{
		Val: 1,
		left: &Node{
			Val: 2,
			left: &Node{
				Val: 4,
			},
			right: &Node{
				Val: 5,
			},
		},

		right: &Node{
			Val: 3,
			left: &Node{
				Val: 6,
			},
			right: &Node{
				Val: 7,
			},
		},
	}
	preOrder(header)
	println()
	inOrder(header)
	println()
	blackOrder(header)
	println()
	val := preorderTraversal(header)
	for _, v := range val {
		println(v)
	}
}

func TestAbc(t *testing.T) {
	println(3 / 2)
}

func TestLevelOrder(t *testing.T) {
	levelOrder()
}
