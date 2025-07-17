package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//数组数量
	p1, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	strconv.Atoi(scanner.Text())

	scanner.Scan()
	p3 := scanner.Text()

	scanner.Scan()
	p4, _ := strconv.Atoi(scanner.Text())

	buzhong := strings.Fields(p3)
	ab := map[int]bool{}
	for _, v := range buzhong {
		r, _ := strconv.Atoi(v)
		ab[r-1] = true
	}

	nums := make([]int, p1)
	for i := 0; i < p1; i++ {
		if ab[i] {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	fmt.Println(nums)
	ans := 0
	left, lsum, rsum := 0, 0, 0
	for right, v := range nums {
		rsum += 1 - v
		for rsum-lsum > p4 {
			lsum += 1 - nums[left]
			left++
		}
		ans = max2(ans, right-left+1)
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
