package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main5() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, n)
	dp := make([][]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		arr[i] = v
	}
	//-1标识为计算过
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}
	var f func(L, R int) int
	f = func(L, R int) int {
		if dp[L][R] != -1 {
			return dp[L][R]
		}
		//模拟对手选择方案
		if arr[L] > arr[R] {
			L = (L + 1) % n
		}
		if arr[L] < arr[R] {
			R = (R - 1 + n) % n
		}
		if arr[L] == arr[R] {
			//表示到了最后一位
			dp[L][R] = arr[L]
		} else {
			//选择左边
			leftOptions := arr[L] + f((L+1)%n, R)
			//选择右边
			rightOptions := arr[R] + f(L, (R-1+n)%n)
			if leftOptions > rightOptions {
				dp[L][R] = leftOptions
			} else {
				dp[L][R] = rightOptions
			}
		}
		return dp[L][R]
	}
	for i := 0; i < n; i++ {
		left := (i + 1) % n
		right := (i - 1 + n) % n
		ans = max2(ans, f(left, right)+arr[i])
	}

	fmt.Println(ans)
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
