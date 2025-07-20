package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
题目描述
部门准备举办一场王者荣耀表演赛，有 10 名游戏爱好者参与，分为两队，每队 5 人。

每位参与者都有一个评分，代表着他的游戏水平。为了表演赛尽可能精彩，我们需要把 10 名参赛者分为示例尽量相近的两队。

一队的实力可以表示为这一队 5 名队员的评分总和。

现在给你 10 名参与者的游戏水平评分，请你根据上述要求分队，最后输出这两组的实力差绝对值。

例：10 名参赛者的评分分别为：5 1 8 3 4 6 7 10 9 2，分组为（1 3 5 8 10）和（2 4 6 7 9），两组实力差最小，差值为1。有多种分法，但是实力差的绝对值最小为1。

输入描述
10个整数，表示10名参与者的游戏水平评分。范围在 [1, 10000] 之间。

输出描述
1个整数，表示分组后两组实力差绝对值的最小值。
*/

/*
*
输入：1 2 3 4 5 6 7 8 9 10
输出：1
*/
func main18() {
	//转化为背包问题：dp[j]表示背包的容量，该题转化为 j=总人数/2，然后10-2*j
	//背包推导公式：dp[j] = max(dp[j], dp[j-score[i]]+score[i])
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	arr := make([]int, len(input))
	count := 0
	for i := range arr {
		arr[i], _ = strconv.Atoi(input[i])
		count += arr[i]
	}
	//目标
	target := count / 2
	//转化为背包问题：一位数组：dp[j]=max(dp[j],dp[j-score[i]]+value[i])
	//先遍历物品，在遍历背包，其背包要倒叙遍历
	dp := make([]int, count)
	for i := 0; i < 10; i++ { //物品
		for j := target; j >= arr[i]; j-- { //背包
			dp[j] = max2(dp[j], dp[j-arr[i]]+arr[i])
		}
	}
	fmt.Println(count - 2*dp[target])
}
