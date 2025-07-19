/*
*
题目描述
给定一个 m x n 的矩阵，由若干字符 ‘X’ 和 ‘O’构成，’X’表示该处已被占据，’O’表示该处空闲，请找到最大的单入口空闲区域。

解释：

空闲区域是由连通的’O’组成的区域，位于边界的’O’可以构成入口，

单入口空闲区域即有且只有一个位于边界的’O’作为入口的由连通的’O’组成的区域。
如果两个元素在水平或垂直方向相邻，则称它们是“连通”的。

输入描述
第一行输入为两个数字，第一个数字为行数m，第二个数字为列数n，两个数字以空格分隔，1<=m,n<=200。

剩余各行为矩阵各行元素，元素为‘X’或‘O’，各元素间以空格分隔。

输出描述
若有唯一符合要求的最大单入口空闲区域，输出三个数字

第一个数字为入口行坐标（0~m-1）
第二个数字为入口列坐标（0~n-1）
第三个数字为区域大小
*/

/*
输入：
5 4
X X X X
X O O O
X X X X
X O O O
45
输出：
3
*/
/**
思路：使用递归深度搜索，查询i、j的向上、向下、向左、向右的遍历查询
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	rows, cols int
	matrix     [][]string
	// 方向偏移量：上、下、左、右
	offsets = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	// 记录已访问的位置
	visited map[string]bool
)

func main12() {
	// 读取输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), " ")
	rows, _ = strconv.Atoi(firstLine[0])
	cols, _ = strconv.Atoi(firstLine[1])

	// 初始化矩阵
	matrix = make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	// 填充矩阵数据
	for i := 0; i < rows; i++ {
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")
		for j := 0; j < cols; j++ {
			matrix[i][j] = line[j]
		}
	}

	// 计算并输出结果
	fmt.Println(getResult())
}

func getResult() string {
	// 存储符合条件的区域：每个区域包含入口坐标和区域大小
	var regions [][]int
	// 初始化已访问记录
	visited = make(map[string]bool)

	// 遍历矩阵的每个位置
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 遇到未访问的'O'，开始DFS遍历
			if matrix[i][j] == "O" && !visited[fmt.Sprintf("%d-%d", i, j)] {
				// 存储当前区域的边界入口
				var entrances [][]int
				// 执行DFS，返回区域大小
				size := dfs(i, j, &entrances)

				// 如果只有一个入口，则记录该区域
				if len(entrances) == 1 {
					entrance := entrances[0]
					regions = append(regions, []int{entrance[0], entrance[1], size})
				}
			}
		}
	}

	// 没有符合条件的区域
	if len(regions) == 0 {
		return "NULL"
	}

	// 按区域大小降序排序
	sort.Slice(regions, func(i, j int) bool {
		return regions[i][2] > regions[j][2]
	})

	// 检查最大区域是否唯一
	maxSize := regions[0][2]
	if len(regions) > 1 && regions[1][2] == maxSize {
		return strconv.Itoa(maxSize)
	}

	// 返回最大区域的入口坐标和大小
	return fmt.Sprintf("%d %d %d", regions[0][0], regions[0][1], regions[0][2])
}

func dfs(i, j int, entrances *[][]int) int {
	// 检查位置是否越界
	if i < 0 || i >= rows || j < 0 || j >= cols {
		return 0
	}

	// 生成位置标识符
	pos := fmt.Sprintf("%d-%d", i, j)

	// 检查是否已访问或是障碍物
	if matrix[i][j] == "X" || visited[pos] {
		return 0
	}

	// 标记当前位置为已访问
	visited[pos] = true

	// 如果当前位置在边界上，则记录为入口
	if i == 0 || i == rows-1 || j == 0 || j == cols-1 {
		*entrances = append(*entrances, []int{i, j})
	}

	// 当前单元格计数
	count := 1

	// 向四个方向进行DFS遍历
	for _, offset := range offsets {
		newI := i + offset[0]
		newJ := j + offset[1]
		count += dfs(newI, newJ, entrances)
	}

	return count
}
