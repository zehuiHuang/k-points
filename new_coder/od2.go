package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 返回矩阵中非1的元素个数

/*
*
题目描述
存在一个m*n的二维数组，其成员取值范围为0，1，2。

其中值为1的元素具备同化特性，每经过1S，将上下左右值为0的元素同化为1。

而值为2的元素，免疫同化。

将数组所有成员随机初始化为0或2，再将矩阵的[0, 0]元素修改成1，在经过足够长的时间后求矩阵中有多少个元素是0或2（即0和2数量之和）。

输入描述
输入的前两个数字是矩阵大小。后面是数字矩阵内容。

输出描述
返回矩阵中非1的元素个数。
*/

/*
*
输入	：
4 4
0 0 0 0
0 2 2 2
0 2 0 0
0 2 0 0
输出	：
9
*/
func main2() {
	offset := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	//思路：深度优先,从第一个位置，同时对该位置进行上下左右的扩散
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mn := strings.Fields(scanner.Text())
	m, _ := strconv.Atoi(mn[0])
	n, _ := strconv.Atoi(mn[1])

	tables := make([][]string, m)
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		tables[i] = strings.Fields(scanner.Text())
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		current := tables[i][j]
		if current != "0" {
			return
		}
		tables[i][j] = "1"
		visited[i][j] = true
		for k := range offset {
			newX := i + offset[k][0]
			newY := j + offset[k][1]
			if newX < m && newX >= 0 && newY < n && newY >= 0 && tables[newX][newY] == "0" && !visited[newX][newY] {
				//递归
				dfs(newX, newY)
			}
		}
	}
	tables[0][0] = "0"
	dfs(0, 0)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tables[i][j] != "1" {
				count++
			}
		}
	}
	fmt.Println(count)

}
