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
题目描述
存在一个m×n的二维数组，其成员取值范围为0或1。

其中值为1的成员具备扩散性，每经过1S，将上下左右值为0的成员同化为1。

二维数组的成员初始值都为0，将第[i,j]和[k,l]两个个位置上元素修改成1后，求矩阵的所有元素变为1需要多长时间。

输入描述
输入数据中的前2个数字表示这是一个m×n的矩阵，m和n不会超过1024大小；

中间两个数字表示一个初始扩散点位置为i,j；

最后2个数字表示另一个扩散点位置为k,l。

输出描述
输出矩阵的所有元素变为1所需要秒数
*/

/*
*
输入：
4,4,0,0,3,3
输出：3
*/
/**
思路：将每次即将扩散的点放入队列。循环队列中的坐标，之后做两件事情：1是四周扩散找到坐标并感染（设置为1），
之后将感染后的点放入新的队列以便重新进行下一轮的感染，等待第一轮的队列遍历完后，开始遍历新的队列
*/
func main13() {

	offsets := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	scannner := bufio.NewScanner(os.Stdin)
	scannner.Scan()
	arr := strings.Split(scannner.Text(), ",")
	m, _ := strconv.Atoi(arr[0])
	n, _ := strconv.Atoi(arr[1])
	x1, _ := strconv.Atoi(arr[2])
	y1, _ := strconv.Atoi(arr[3])

	x2, _ := strconv.Atoi(arr[4])
	y2, _ := strconv.Atoi(arr[5])

	table := make([][]int, m)
	for i, _ := range table {
		table[i] = make([]int, n)
	}
	table[x1][y1] = 1
	table[x2][y2] = 1
	queue := make([][]int, 0)
	queue = append(queue, []int{x1, y1}, []int{x2, y2})
	time := 0
	count := m*n - 2
	for len(queue) > 0 && count > 0 {
		queueNew := make([][]int, 0)
		//遍历队列中的所有数据坐标
		for i := range queue {
			x := queue[i][0]
			y := queue[i][1]
			for k := 0; k < len(offsets); k++ {
				newx := x + offsets[k][0]
				newy := y + offsets[k][1]
				//符合扩散的条件
				if newx >= 0 && newx < m && newy >= 0 && newy < n && table[newx][newy] == 0 {
					table[newx][newy] = 1
					queueNew = append(queueNew, []int{newx, newy})
					count--
				}
			}
		}
		queue = queueNew
		time++
	}
	fmt.Println(time)
}
