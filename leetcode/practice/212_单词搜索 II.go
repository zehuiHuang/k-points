package practice

/*
*
思路,针对每个单词进行处理,和单词搜索的区别:
1、单词变成了多个
2、要返回具体的符合条件的单词
*/

// 以下算法虽然逻辑正确,但时间超限
func findWords(board [][]byte, words []string) []string {
	ans := []string{}
	if len(board) == 0 || len(board[0]) == 0 {
		return ans
	}
	offset := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(board), len(board[0])

	var dfs func(i, j, k int, word string) bool
	dfs = func(i, j, k int, word string) bool {

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
			if dfs(newX, newY, k+1, word) {
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
			for _, v := range words {
				if dfs(i, j, 0, v) {
					ans = append(ans, v)
				}
			}
		}
	}
	return ans
}

func findWords2(board [][]byte, words []string) []string {
	ans := []string{}
	return ans
}
