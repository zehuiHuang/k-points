package practice

/*
*
思路: BFS

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
*/
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0
	queue := [][2]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//腐烂的橘子
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				//统计新鲜橘子,
				count++
			}
		}
	}
	time := 0
	for count > 0 && len(queue) > 0 {
		time++
		length := len(queue)
		for i := 0; i < length; i++ {
			p := queue[0]
			queue = queue[1:]
			r := p[0] //行坐标
			c := p[1] //列坐标
			//新鲜的橘子
			if r-1 >= 0 && r-1 < m && grid[r-1][c] == 1 {
				grid[r-1][c] = 2 //腐烂
				count--
				queue = append(queue, [2]int{r - 1, c})
			}

			if r+1 < m && grid[r+1][c] == 1 {
				grid[r+1][c] = 2 //腐烂
				count--
				queue = append(queue, [2]int{r + 1, c})
			}

			if c-1 >= 0 && c-1 < n && grid[r][c-1] == 1 {
				grid[r][c-1] = 2 //腐烂
				count--
				queue = append(queue, [2]int{r, c - 1})
			}

			if c+1 < n && grid[r][c+1] == 1 {
				grid[r][c+1] = 2 //腐烂
				count--
				queue = append(queue, [2]int{r, c + 1})
			}
		}
	}
	if count > 0 {
		return -1
	}
	return time
}
