package practice

import "math"

/**
思路：每次循环找一次历史最低价格，然后将历史最高价格减去历史最低价，即可得到相差的最大值
*/

func maxProfit(prices []int) int {
	n := len(prices)
	//最小值
	maxP := 0
	//最大差值
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
