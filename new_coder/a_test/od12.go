package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
输入：
4 5
X X X X X
O O O O X
X O O O X
X O X X O
输出：
3 4 1

输入：
5 4
X X X X
X O O O
X O O O
X O O X
X X X X
输出：NULL

输入：
5 4
X X X X
X O O O
X X X X
X O O O
X X X X
输出：3
*/

func main() {
	offsets := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a := strings.Fields(scanner.Text())
	m, _ := strconv.Atoi(a[0]) //行
	n, _ := strconv.Atoi(a[1]) //列
	visited := make(map[string]bool)
	arr := make([][]string, m)
	for i, _ := range arr {
		scanner.Scan()
		h := strings.Fields(scanner.Text())
		arr[i] = h
	}
	fmt.Println(arr)
	fmt.Println("---------")

	var dfs func(i, j int, entrance *[][]int) int

	dfs = func(i, j int, entrance *[][]int) int {
		//如果超过坐标范围
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		key := fmt.Sprintf("%d_%d", i, j)
		if arr[i][j] == "X" || visited[key] {
			return 0
		}
		if i == 0 || i == m-1 || j == 0 || j == n-1 {
			*entrance = append(*entrance, []int{i, j})
		}
		visited[key] = true
		count := 1
		//向左向右向上向下进行递归计算count
		for k := 0; k < len(offsets); k++ {
			count += dfs(i+offsets[k][0], j+offsets[k][1], entrance)
		}
		return count
	}
	//结果集合
	ans := make([][]int, 0)
	//m、n、arr
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			entrance := make([][]int, 0)
			key := fmt.Sprintf("%d_%d", i, j)
			if arr[i][j] == "O" && !visited[key] {
				count := dfs(i, j, &entrance)
				if len(entrance) == 1 {
					ans = append(ans, []int{entrance[0][0], entrance[0][1], count})
				}
			}
		}
	}
	if len(ans) == 0 {
		fmt.Println("NULL")
		return
	}
	//对ans 按照count进行降序排序
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][2] > ans[j][2]
	})
	if len(ans) > 1 && ans[0][2] == ans[1][2] {
		fmt.Println(ans[0][2])
	} else {
		fmt.Println(ans[0])
	}
}
