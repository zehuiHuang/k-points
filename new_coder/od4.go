package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//传递悄悄话
/*
*
给定一个二叉树，每个节点上站一个人，节点数字表示父节点到该节点传递悄悄话需要花费的时间。

初始时，根节点所在位置的人有一个悄悄话想要传递给其他人，求二叉树所有节点上的人都接收到悄悄话花费的时间。

*/

/*
*输入：
0 9 20 -1 -1 15 7 -1 -1 -1 -1 3 2
输出：
38
*/
func main4() {

	//思路：层序遍历，并且利用数组和链表的转换逻辑
	//在层序遍历时，将父节点的消耗时间加到子节点当种

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
	ansTemp := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			//索引值
			qi := queue[0]
			//索引值对应的消耗时间（父节点到当前节点的消息传递时间）
			qv := arr[qi]
			queue = queue[1:]
			//父节点的左子节点
			leftNodeIndex := 2*qi + 1
			if leftNodeIndex < len(arr) && arr[leftNodeIndex] != -1 {
				arr[leftNodeIndex] += qv
				queue = append(queue, leftNodeIndex)
				ansTemp = max2(ansTemp, arr[leftNodeIndex])
			}
			//父节点的右子节点
			rightNodeIndex := 2*qi + 2
			if rightNodeIndex < len(arr) && arr[rightNodeIndex] != -1 {
				arr[rightNodeIndex] += qv
				queue = append(queue, rightNodeIndex)
				ansTemp = max2(ansTemp, arr[leftNodeIndex])
			}
		}
	}
	fmt.Println(ansTemp)
}
