package practice

/*
*
思路:回溯
遍历二维所有位置,然后以该位置为起点,向四周扩散,并收集结果,如果结果符合条件则直接返回即可
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	offset := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(board), len(board[0])

	var dfs func(board [][]byte, i, j, k int) bool
	dfs = func(board [][]byte, i, j, k int) bool {

		// 先检查边界
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		// 检查当前字符是否匹配
		if board[i][j] != word[k] {
			return false
		}

		//满足条件
		if k == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = '0'
		//从四个方向进行扩撒
		for index := range offset {
			newX := i + offset[index][0]
			newY := j + offset[index][1]
			//终止条件
			if dfs(board, newX, newY, k+1) {
				//回溯,重置
				board[i][j] = temp
				return true
			}
		}
		board[i][j] = temp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res := dfs(board, i, j, 0)
			if res {
				return true
			}
		}
	}
	return false
}
