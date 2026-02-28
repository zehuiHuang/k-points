package practice

/*
*
LeetCode 139. 单词拆分
https://leetcode.cn/problems/word-break/

思路：动态规划
dp[i] 表示字符串 s 的前 i 个字符（s[0:i-1]）能否被拆分成字典中的单词
状态转移方程：dp[i] = dp[j] && wordDictSet.contains(s[j:i-1])，其中 0 <= j < i
*/
func wordBreak(s string, wordDict []string) bool {
	// 将单词字典转换为哈希集合，方便快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true // 空字符串可以被拆分

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 如果前 j 个字符可以被拆分，并且 s[j:i] 在字典中
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // 找到一个有效的拆分即可
			}
		}
	}

	return dp[n]
}

/*
*
优化版本：使用字典中最长单词长度进行剪枝
*/
func wordBreakOptimized(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	maxLen := 0
	for _, word := range wordDict {
		wordSet[word] = true
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		// j 从 i-maxLen 开始，但至少为 0
		start := i - maxLen
		if start < 0 {
			start = 0
		}
		for j := start; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
