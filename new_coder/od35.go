package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 欢乐的周末

/*
*
题目描述
小华和小为是很要好的朋友，他们约定周末一起吃饭。

通过手机交流，他们在地图上选择了多个聚餐地点（由于自然地形等原因，部分聚餐地点不可达），求小华和小为都能到达的聚餐地点有多少个？

输入描述
第一行输入 m 和 n

m 代表地图的长度
n 代表地图的宽度
第二行开始具体输入地图信息，地图信息包含：

0 为通畅的道路
1 为障碍物（且仅1为障碍物）
2 为小华或者小为，地图中必定有且仅有2个 （非障碍物）
3 为被选中的聚餐地点（非障碍物）
输出描述
可以被两方都到达的聚餐地点数量，行末无空格。
*/
/**
输入
4 4
2 1 0 3
0 1 2 1
0 3 0 0
0 0 0 0
输出
2
*/
func main35() {
	//思路：根据并查集(长度为m*n)来计算小华小为是否在一个集合，如果在一个集合，在计算哪些聚餐的地和小伟（或小华）属于同一个集合
	//1、通过遍历二维数组中的每一个位置，统计出聚餐地、小华小为所在地，
	//2、对非1的点进行向左、向右、向下、向上进行扩散，然后进行集合合并
	scanenr := bufio.NewScanner(os.Stdin)
	scanenr.Scan()
	mn := strings.Fields(scanenr.Text())
	m, _ := strconv.Atoi(mn[0])
	n, _ := strconv.Atoi(mn[1])

	//组装二维数组
	tables := make([][]int, m)
	for i := range tables {
		tables[i] = make([]int, n)
		scanenr.Scan()
		rs := strings.Fields(scanenr.Text())
		for j := range rs {
			tables[i][j], _ = strconv.Atoi(rs[j])
		}
	}
	offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	//小华、小为所在位置
	p := make([]int, 0)
	//聚餐地位置
	f := make([]int, 0)

	//并查集和初始化
	unions := make([]int, m*n)
	for i := range unions {
		unions[i] = i
	}
	var find func(parent []int, x int) int
	find = func(parent []int, x int) int {
		if parent[x] != x {
			parent[x] = find(parent, parent[x])
		}
		return parent[x]
	}
	var union func(parent []int, x, y int)
	union = func(parent []int, x, y int) {
		xx := find(parent, x)
		yy := find(parent, y)
		if xx != yy {
			parent[xx] = yy
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tables[i][j] != 1 {
				pos := i*n + j //转为一位数组
				//小华、小为的坐标
				if tables[i][j] == 2 {
					p = append(p, pos)
				} else if tables[i][j] == 3 {
					f = append(f, pos)
				}
				for k := range offsets {
					newX := i + offsets[k][0]
					newY := j + offsets[k][1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n && tables[newX][newY] != 1 {
						union(unions, pos, newX*m+newY)
					}
				}
			}
		}
	}
	h := find(unions, p[0])
	w := find(unions, p[1])
	if h != w {
		fmt.Println(0)
		return
	}
	ans := 0
	for i := range f {
		fd := find(unions, f[i])
		if fd == h {
			ans++
		}
	}
	fmt.Println(ans)
}
