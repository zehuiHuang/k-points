package leetcode

/**
最短路径问题
*/

// findCheapestPrice 找到从src到dst的最便宜路径，最多经过k次中转
// 参数：
//
//	n: 城市数量
//	flights: 航班信息，每个元素是[j, i, cost]，表示从j到i的航班费用为cost
//	src: 起始城市
//	dst: 目标城市
//	k: 最大中转次数
//
// 返回：最便宜的机票价格，如果无法到达则返回-1
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 10000*101 + 1
	f := make([][]int, k+2)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	f[0][src] = 0
	for t := 1; t <= k+1; t++ {
		for _, flight := range flights {
			j, i, cost := flight[0], flight[1], flight[2]
			f[t][i] = min(f[t][i], f[t-1][j]+cost)
		}
	}
	ans := inf
	for t := 1; t <= k+1; t++ {
		ans = min(ans, f[t][dst])
	}
	if ans == inf {
		ans = -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
