package backtrack

import (
	"strconv"
	"strings"
)

/**
回溯类算法---------------------------------------------------
*/

// 93. 复原 IP 地址
// 思路：回溯算法，通过一个个决策树来切割
func restoreIpAddresses(s string) []string {
	//参数为字符串，返回值的第一个参数为搜索树每一层的选择结果，第二个参数为符合结果的集合
	res := []string{}
	var dfs func(s string, tpms []string)
	dfs = func(s string, tpms []string) {
		if len(tpms) == 4 && s != "" {
			return
		}
		if s == "" && len(tpms) > 4 {
			return
		}
		if s == "" && len(tpms) == 4 {
			tmp := tpms[0]
			for i := 1; i < len(tpms); i++ {
				tmp += "." + tpms[i]
			}
			res = append(res, tmp)
			return
		}

		for i := 1; i <= 3 && i <= len(s); i++ {
			subStr := s[:i]
			// 前导零检查 (排除"00", "01"等)
			if len(subStr) > 1 && subStr[0] == '0' {
				break // 后续长度也无效，直接退出
			}
			num, _ := strconv.Atoi(subStr)
			if num > 255 {
				break // 后续长度数字更大，提前退出
			}
			// 递归回溯
			dfs(s[i:], append(tpms, subStr))
		}

	}
	tmps := []string{}
	dfs(s, tmps)
	return res
}

// 131. 分割回文串
// 思路：回溯算法，通过决策树进行切割
func partition(s string) [][]string {
	//s = "aab"
	//切割a\ab-》留a，在对ab切割
	ans := [][]string{}
	tmps := []string{}
	var dfs func(s string, tmps []string)
	dfs = func(s string, tmps []string) {
		//判断截止条件
		if len(s) == 0 {
			pathCopy := make([]string, len(tmps))
			copy(pathCopy, tmps)
			ans = append(ans, pathCopy)
			return
		}
		for i := 1; i <= len(s); i++ {
			tmp := s[:i]
			//判断tmp是否是回文串
			if isPalindrome(tmp) {
				dfs(s[i:], append(tmps, tmp))
			}
		}
	}
	dfs(s, tmps)
	return ans
}

// 判断str是否是回文串
func isPalindrome(str string) bool {
	start, end := 0, len(str)-1
	for start < end {
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// N皇后问题
// 思路：回溯方法：每一行进行尝试，在每一行尝试下的每一列进行尝试（类似切割），然后将符合条件的拼装
func solveNQueens(n int) [][]string {
	//对棋盘进行初始化
	ans := [][]string{}
	qipan := make([][]string, n)
	for i := 0; i < n; i++ {
		qipan[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			qipan[i][j] = "."
		}
	}
	//temps存储临时数据
	var dfs func(row int, temps []string)
	dfs = func(row int, temps []string) {
		if row == n {
			//搜集答案
			pathCopy := make([]string, len(temps))
			copy(pathCopy, temps)
			ans = append(ans, pathCopy)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, qipan) {
				qipan[row][i] = "Q"
				dfs(row+1, append(temps, strings.Join(qipan[row], "")))
				qipan[row][i] = "."
			}
		}
	}
	temps := []string{}
	dfs(0, temps)
	return ans
}

func isValid(n, row, col int, chessboard [][]string) bool {
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	return true
}

// 216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	var dfs func(startIndex, sum int, temps []int)
	dfs = func(startIndex, sum int, temps []int) {
		if len(temps) == k && sum == n {
			newTemps := make([]int, k)
			copy(newTemps, temps)
			ans = append(ans, newTemps)
			return
		}
		for i := startIndex; i <= 9; i++ {
			dfs(i+1, sum+i, append(temps, i))
		}
	}
	dfs(1, 0, []int{})
	return ans
}

// 77. 组合
// n=1,2,3,4
// k=2
func combine(n int, k int) [][]int {
	//结果
	ans := [][]int{}
	var dfs func(startIndex int, path []int)
	dfs = func(startIndex int, path []int) {
		if len(path) == k {
			newTemps := make([]int, len(path))
			copy(newTemps, path)
			ans = append(ans, newTemps)
			return
		}
		//重点理解:for循环是横向遍历,递归是纵向遍历
		//横向遍历是为了从1开始取值,比如取1,取2
		//纵向遍历是为了在取1的基础上,再取下一个,组合并判断是否符合条件
		for i := startIndex; i <= n; i++ {
			dfs(i+1, append(path, i))
		}

	}
	dfs(1, []int{})
	return ans
}
