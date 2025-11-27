package practice

/**
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

示例 1：

输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
示例 2：

输入：s = "a"
输出：[["a"]]

*/

// 回溯:进行抽象树模拟

func partition(s string) [][]string {
	//
	ans := [][]string{}
	var dfs func(s string, temp []string)
	dfs = func(s string, temp []string) {
		//截止条件,并收集结果
		if len(s) == 0 {
			tmp := make([]string, len(temp))
			copy(tmp, temp)
			ans = append(ans, tmp)
			return
		}
		//切割
		for i := 1; i <= len(s); i++ {
			t := s[:i]
			//如果是回文,则继续对剩下的继续切割
			if isPalindrome(t) {
				dfs(s[i:], append(temp, t))
			}
		}
	}
	dfs(s, []string{})
	return ans
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
