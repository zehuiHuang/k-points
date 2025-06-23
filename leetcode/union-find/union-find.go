package union_find

/*
*
leetcode 547 省份数量
https://leetcode.cn/problems/bLyHh0/solutions/1087069/sheng-fen-shu-liang-by-leetcode-solution-c8b8/
并查集
输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
输出：2

1,1,0
1,1,0
0,0,1

关于城市数量的理解，下标索引值即为城市的名字或标识，n*n 最多就有n个城市，比如x坐标：0，1，2 y坐标：0，1，2
*/
func findCircleNum(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)
	//父节点初始值为下标值

	//value：0 1 2 ... n
	//index：0 1 2 ... n
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	//查询x的根节点(也是下标)
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	//合并两个城市，将from集和 合并 到to集和中
	//0  1  2 3  4  5  6 7 8 9 父节下标
	//0  1  2 3  4  5  6 7 8 9 节点值
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		//row，表示第i个城市的连接情况
		for j := i + 1; j < n; j++ {
			//row[j]表示第i个城市和第j个城市是否相连，如果相连，则将两个城市合并（将i合并到j集和中，也就是将下标为i的值设置为j根节点的下标）
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	//当数组中下标和值相等则即为根节点
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}

// leetcode 695. 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	//m:行长度，n:列长度
	m, n := len(grid), len(grid[0])

	// 初始化并查集（父节点数组 + 面积数组）
	parent := make([]int, m*n)
	area := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := i*n + j // 二维转一维：为了更方便地用一维数组（或列表）来管理网格中的每个单元格
			parent[index] = index
			if grid[i][j] == 1 {
				area[index] = 1 // 初始陆地面积为1
			}
		}
	}

	// 定义方向数组（向右、向下）
	dirs := [][]int{{0, 1}, {1, 0}}

	// 遍历网格合并相邻陆地
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			// 检查右方和下方（避免重复处理）
			for _, dir := range dirs {
				ni, nj := i+dir[0], j+dir[1]
				if ni < m && nj < n && grid[ni][nj] == 1 {
					union(parent, area, i*n+j, ni*n+nj)
				}
			}
		}
	}

	// 查找最大面积
	maxArea := 0
	for i := range area {
		if area[i] > maxArea {
			maxArea = area[i]
		}
	}
	return maxArea
}

// 带路径压缩的查找
func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// 合并并维护面积
func union(parent, area []int, x, y int) {
	rootX := find(parent, x)
	rootY := find(parent, y)
	if rootX == rootY {
		return
	}
	// 合并到面积较大的根（优化方向可选）
	if area[rootX] < area[rootY] {
		parent[rootX] = rootY
		area[rootY] += area[rootX]
	} else {
		parent[rootY] = rootX
		area[rootX] += area[rootY]
	}
}
