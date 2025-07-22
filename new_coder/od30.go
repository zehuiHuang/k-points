package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//开心消消乐
/*
*
给定一个 N 行 M 列的二维矩阵，矩阵中每个位置的数字取值为 0 或 1。
矩阵示例如：1 1 0 00 0 0 10 0 1 11 1 1 1现需要将矩阵中所有的 1 进行反转为 0，
规则如下：当点击一个 1 时，该 1 便被反转为0，同时相邻的上、下、左、右，以及左上、左下、右上、右下 8 个方向的 1（如果存在1）均会自动反转为 0
进一步地，一个位置上的 1 被反转为0时，与其相邻的 8 个方向的 1（如果存在1）均会自动反转为0
按照上述规则示例中的矩阵只最少需要点击 2 次后，所有值均为 0。请问，给定一个矩阵，最少需要点击几次后，所有数字均为 0？
*/

/*
*输入：
4 4
1 1 0 0
0 0 0 1
0 0 1 1
1 1 1 1
输出：
2
*/
func main30() {
	defi := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mn := strings.Fields(scanner.Text())
	m, _ := strconv.Atoi(mn[0])
	n, _ := strconv.Atoi(mn[1])

	table := make([][]int, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		rows := strings.Fields(scanner.Text())
		table[i] = make([]int, n)
		for j := 0; j < n; j++ {
			table[i][j], _ = strconv.Atoi(rows[j])
		}
	}

	//思路:
	valid := make([][]bool, m)
	for i := 0; i < m; i++ {
		valid[i] = make([]bool, n)
	}

	var bfs func(table [][]int, m, n int) int
	clickCount := 0
	bfs = func(table [][]int, m, n int) int {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if table[i][j] == 1 && !valid[i][j] {
					clickCount++
					queue := [][2]int{{i, j}}
					valid[i][j] = true
					for len(queue) > 0 {
						point := queue[0]
						queue = queue[1:]
						x := point[0]
						y := point[1]
						for k := 0; k < len(defi); k++ {
							newX := x + defi[k][0]
							newY := y + defi[k][1]
							if newX < m && newX >= 0 && newY < n && newY >= 0 && table[newX][newY] == 1 && !valid[newX][newY] {
								valid[newX][newY] = true
								queue = append(queue, [2]int{newX, newY})
							}
						}
					}
				}
			}
		}
		return clickCount
	}
	fmt.Println(bfs(table, m, n))
}
