package practice

/*
*
思路:将两个字符串的索引对应成一个二位数组(word1的索引位置是m,word2的索引位置是n)
dp[i][j] 代表 源字符串word1前i个字符，变成与目标字符串word2前j个字符一模一样需要的编辑次数。

1、其中的任意i，j位置，如果两个字符串的相同索引位置对应的字符不同,则都和他的上方、左方、左上方有关系，取这三个方向的最小值
即:dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
2、如果相同则dp[i][j]=dp[i-1][j-1]
*/

/*
*
理解上方、左方、左上方
左上方:
dp[i-1][j-1](表示替换)：word1 的前 i-1 个字符已经能匹配 word2 的前 j-1 个字符，现在只需要在 word1 末尾替换一个字符（即 word2 的第 j 个字符），就能匹配 word2 的前 j 个字符。
上方:
dp[i][j-1] (表示插入)：word1 的前 i 个字符已经能匹配 word2 的前 j-1 个字符，现在只需要在 word1 末尾插入一个字符（即 word2 的第 j 个字符），就能匹配 word2 的前 j 个字符。
左方:
dp[i-1][j] (表示删除)：word1 的前 i-1 个字符已经能匹配 word2 的前 j 个字符，现在只需要在 word1 末尾删除一个字符（即 word1 的第 i 个字符），就能匹配 word2 的前 j 个字符。
*/
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//第一行的每一列
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	//第一列的每一行
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
