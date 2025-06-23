package greedy

import (
	"container/heap"
	"math"
)

//贪心算法

// https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0134.%E5%8A%A0%E6%B2%B9%E7%AB%99.md
/*
*
加油站：leetcode:134 普通方法
*/
func canCompleteCircuit(gas, cost []int) int {
	N := len(gas)
	for i := 0; i < N; i++ {
		//当前的第几个结点，目的是为了判断是否走完一圈
		index := (i + 1) / N
		//到达index结点的剩余油量
		remain := gas[i] - cost[i]
		for remain > 0 && index != i {
			remain += gas[index] - cost[index]
			index = (index + 1) % N
		}
		if remain >= 0 && index == i {
			return index
		}
	}
	return -1
}

/*
*
加油站：leetcode:134 贪心算法
*/
func canCompleteCircuit2(gas, cost []int) int {
	n := len(gas)
	//起点
	start := 0
	//一路总的剩余油量
	totalCount := 0
	//从起点开始的当前剩余油量
	currentCount := 0
	for i := 0; i < n; i++ {
		totalCount += gas[i] - cost[i]
		currentCount += gas[i] - cost[i]
		//如果当前剩余油量小于0，说明从start到i这段路无法走完，所以从i+1开始重新计算
		if currentCount < 0 {
			start = i + 1
			currentCount = 0
		}
	}
	//如果总的剩余油量小于0，说明无法走完一圈
	if totalCount < 0 {
		return -1
	}
	return start
}

/*
买卖股票的最佳时机：leetcode:121 贪心算法

200 300 100 500 200
0   1   2   3   4
思路：每次循环找一次历史最低价格，然后将历史最高价格减去历史最低价，即可得到相差的最大值
*/

func maxProfit(prices []int) int {
	n := len(prices)
	maxP := 0
	minPrice := math.MaxInt64
	for i := 0; i < n; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxP {
			maxP = prices[i] - minPrice
		}
	}
	return maxP
}

/*
买卖股票的最佳时机：leetcode:122 贪心算法

7 1 5 3 6 4
0 1 2 3 4`5
思路：画一个折线图，由于每天都能买卖，那么只计算有收益的就行，然后把每天的收益相加即可
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	sum := 0
	for i := 1; i < n; i++ {
		if prices[i+1] > prices[i] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}

/*
*
leetcode 3123 最短路径中的边
*/
func findAnswer(n int, edges [][]int) []bool {
	type edge struct{ to, w, i int }
	g := make([][]edge, n)
	for i, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w, i})
		g[y] = append(g[y], edge{x, w, i})
	}

	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x := p.x
		if p.dis > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := p.dis + e.w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}

	ans := make([]bool, len(edges))
	// 图不连通
	if dis[n-1] == math.MaxInt {
		return ans
	}

	// 从终点出发 BFS
	vis := make([]bool, n)
	vis[n-1] = true
	q := []int{n - 1}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		for _, e := range g[y] {
			x := e.to
			if dis[x]+e.w != dis[y] {
				continue
			}
			ans[e.i] = true
			if !vis[x] {
				vis[x] = true
				q = append(q, x)
			}
		}
	}
	return ans
}

type pair struct{ dis, x int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
