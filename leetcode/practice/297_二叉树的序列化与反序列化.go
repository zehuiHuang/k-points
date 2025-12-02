package practice

import (
	"strconv"
	"strings"
)

/*
*
思路: 采用层序遍历将数据存储起来,重点是存储的值和相对位置都要

1,2,x,x,3,4,x,x,5,x,x

1,

左树:2,x,x

4,x,x,5,x,x
*/
type Codec struct {
}

//func Constructor() Codec {
//
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//使用前序遍历将节点存储起来,当节点的左右节点为nil时,也需要存储(存储为nil)
	if root == nil {
		return "nil"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	return buildTree2(&arr)
}

func buildTree2(list *[]string) *TreeNode {
	//拿到头节点
	Val := (*list)[0]
	//将头节点移除
	*list = (*list)[1:]
	if Val == "nil" {
		return nil
	}
	//创建当前节点
	v, _ := strconv.Atoi(Val)
	root := &TreeNode{Val: v}
	root.Left = buildTree2(list)
	root.Right = buildTree2(list)
	return root
}
