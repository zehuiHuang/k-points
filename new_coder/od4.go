package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*
给定一个二叉树，每个节点上站一个人，节点数字表示父节点到该节点传递悄悄话需要花费的时间。

初始时，根节点所在位置的人有一个悄悄话想要传递给其他人，求二叉树所有节点上的人都接收到悄悄话花费的时间。
//思路：层序遍历，并且数组和链表可以进行转换
*/
func main4s() {
	//0 9 20 -1 -1 15 7 -1 -1 -1 -1 3 2
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	inputStrArr := strings.Fields(inputStr)
	arr := make([]int, len(inputStrArr))
	for i, v := range inputStrArr {
		rv, _ := strconv.Atoi(v)
		arr[i] = rv
	}
	queue := []int{0}
	ans := 0
	for len(queue) > 0 {
		length := len(queue)
		ansTemp := 0
		for i := 0; i < length; i++ {
			qi := queue[0]
			qv := arr[qi]
			queue = queue[1:]
			//父节点的左子节点
			leftNodeIndex := 2*qi + 1
			if leftNodeIndex < len(arr) {
				queue = append(queue, leftNodeIndex)
			}
			//父节点的右子节点
			rightNodeIndex := 2*qi + 2
			if rightNodeIndex < len(arr) {
				queue = append(queue, rightNodeIndex)
			}
			ansTemp = max2(ansTemp, qv)
		}
		ans += ansTemp
	}
	fmt.Println(ans)
}
