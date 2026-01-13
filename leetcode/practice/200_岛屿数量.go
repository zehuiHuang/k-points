package practice

/*
*
思路:扩散
遍历矩阵,遇到1则开始扩散
*/
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	count := 0

	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i >= m || i < 0 || j >= n || j < 0 || grid[i][j] == '0' {
			return
		}
		if grid[i][j] == '2' {
			return
		}
		grid[i][j] = '2'
		//向四周开始扩散
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

/**
触类旁通:类似题-------------------------------------------------------------
*/
//695. 岛屿的最大面积
//实现思路除了这种思路,还有并查集
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		if grid[i][j] == 2 {
			return 0
		}
		grid[i][j] = 2
		//四周扩散,并最终返回累加值,到最后最小也是1
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}

// 463. 岛屿的周长
// 思路:周长的计算可以转化为 矩阵为1的格子与网格边界的周长+与海洋边界的周长
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		//与网格边界的周长
		if i < 0 || i >= m || j < 0 || j >= n {
			return 1
		}
		//与海洋边界的周长
		if grid[i][j] == 0 {
			return 1
		}
		//遇到已访问的格子则返回0:因为陆地格子和周长没关系了
		if grid[i][j] == 2 {
			return 0
		}
		grid[i][j] = 2
		//向四周扩散,最小为0(陆地格子和周长没关系了)
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				//题目限制只有一个岛屿，计算一个即可
				return dfs(i, j)
			}
		}
	}
	return res
}
