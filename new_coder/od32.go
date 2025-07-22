package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 战场锁敌
/**
题目描述
有一个大小是N*M的战场地图，被墙壁 ‘#’ 分隔成大小不同的区域，上下左右四个方向相邻的空地 ‘.’ 属于同一个区域，只有空地上可能存在敌人’E”，

请求出地图上总共有多少区域里的敌人数小于K。

输入描述
第一行输入为N,M,K；

N表示地图的行数，M表示地图的列数， K表示目标敌人数量
N，M<=100
之后为一个NxM大小的字符数组。

输出描述
敌人数小于K的区域数量
*/

/*
*
输入：
3 5 2
..#EE
E.#E.
###..

输出：
1
地图被墙壁分为两个区域，左边区域有1个敌人，右边区域有3个敌人，符合条件的区域数量是1
*/

// 思路：深度优先遍历，遍历二位数组中所有的每个位置，只要该位置不是墙壁也未访问过，则用dfs方法迭代计算
func main32() {
	//上下左右
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mnk := strings.Fields(scanner.Text())
	m, _ := strconv.Atoi(mnk[0])
	n, _ := strconv.Atoi(mnk[1])
	k, _ := strconv.Atoi(mnk[2])

	table := make([][]string, m)
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		table[i] = make([]string, n)
		visit[i] = make([]bool, n)
		scanner.Scan()
		c := scanner.Text()
		for j := 0; j < n; j++ {
			table[i][j] = string(c[j])
		}
	}

	ans := 0

	currentCount := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		visit[i][j] = true
		if table[i][j] == "E" {
			currentCount++
		}
		for _, offset := range offsets {
			newX := i + offset[0]
			newY := j + offset[1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && !visit[newX][newY] && table[newX][newY] != "#" {
				dfs(newX, newY)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if table[i][j] == "#" || visit[i][j] {
				continue
			}
			currentCount = 0
			dfs(i, j)
			if currentCount < k && currentCount > 0 {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}
