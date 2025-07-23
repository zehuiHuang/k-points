package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 补种未成活的胡杨

/*
*题目描述
近些年来，我国防沙治沙取得显著成果。某沙漠新种植N棵胡杨（编号1-N），排成一排。
一个月后，有M棵胡杨未能成活。
现可补种胡杨K棵，请问如何补种（只能补种，不能新种），可以得到最多的连续胡杨树？

输入描述
N 总种植数量，1 <= N <= 100000
M 未成活胡杨数量，M 个空格分隔的数，按编号从小到大排列，1 <= M <= N
K 最多可以补种的数量，0 <= K <= M

输出描述
最多的连续胡杨棵树
*/

/*
*
输入：
10
3
2 4 7
1
输出：
6
*/
func main3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//数组的长度
	lenth, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	strconv.Atoi(scanner.Text())

	scanner.Scan()
	p3 := scanner.Text()

	scanner.Scan()
	//能补种的数量
	k, _ := strconv.Atoi(scanner.Text())
	//待补种的下标集合
	buzhong := strings.Fields(p3)
	//待补种的标识,key:切片的坐标索引，v：true表示需要补种，false表示数种存活
	mp := map[int]bool{}
	for _, v := range buzhong {
		r, _ := strconv.Atoi(v)
		mp[r-1] = true
	}

	nums := make([]int, lenth)
	for i := 0; i < lenth; i++ {
		//数组中1表示存活，0表示需要补种
		if mp[i] {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	ans := 0
	left, lsum, rsum := 0, 0, 0
	//利用前缀和 和 滑动窗口
	for right, v := range nums {
		//累加计算从0到right的所有位置上出现0的个数
		rsum += 1 - v
		//若left~right之间的0的数量>k,则左指针向右滑动一个，同时左指针统 从0到left-1上出现的0的数量
		for rsum-lsum > k {
			lsum += 1 - nums[left]
			left++
		}
		ans = max2(ans, right-left+1)
		//其实是ans = max2(ans, right-(left-1))
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
