package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 贪吃的猴子

/*
*
题目描述
一只贪吃的猴子，来到一个果园，发现许多串香蕉排成一行，每串香蕉上有若干根香蕉。每串香蕉的根数由数组numbers给出。

猴子获取香蕉，每次都只能从行的开头或者末尾获取，并且只能获取N次，求猴子最多能获取多少根香蕉。

输入描述
第一行为数组numbers的长度

第二行为数组numbers的值每个数字通过空格分开

第三行输入为N，表示获取的次数

输出描述
按照题目要求能获取的最大数值
*/

/*
*输入：
7
1 2 2 7 3 6 1
3
输出：
10
*/

func main42() {
	//思路：先计算全部选择左边的，右边的不选,然后逐步缩小左边的值A，同时增加右边的值B，那么max(sum+=B-A)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strconv.Atoi(scanner.Text())

	scanner.Scan()
	nums := strings.Fields(scanner.Text())
	tables := make([]int, len(nums))
	for i := range nums {
		tables[i], _ = strconv.Atoi(nums[i])
	}

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	lSum := 0
	rSum := 0
	for i := 0; i < n; i++ {
		lSum += tables[i]
	}

	total := lSum + rSum

	ans := total
	//如果是全选
	if len(tables) == n {
		fmt.Println(n)
		return
	}
	//左边下标，选择减去左边的最后一位
	l := n - 1
	//右边下标，选择加上一位
	r := len(tables) - 1

	for l >= 0 {
		total += (tables[r] - tables[l])
		ans = max2(ans, total)
		r--
		l--
	}
	fmt.Println(ans)
}
