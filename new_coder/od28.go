package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 事件推送
/**
同一个数轴X上有两个点的集合A={A1, A2, …, Am}和B={B1, B2, …, Bn}，Ai和Bj均为正整数，
A、B已经按照从小到大排好序，A、B均不为空，给定一个距离R(正整数)，
列出同时满足如下条件的所有（Ai, Bj）数对：Ai <= Bj Ai, Bj之间的距离小于等于R
在满足1,2的情况下,每个Ai只需输出距离最近的Bj 输出结果按Ai从小到大的顺序排序
*/

/*
*
理解：目的是从setB集合中找到与setA结合的每个元素，保证setB[i]> setA的value[i]，并且差值小于等于r
*/

/*
*
输入：
4 5 5
1 5 5 10
1 3 8 8 20
输出：
1 1
5 8
5 8
*/
func main28() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	//m, _ := strconv.Atoi(input[0])
	n, _ := strconv.Atoi(input[1])
	r, _ := strconv.Atoi(input[2])

	scanner.Scan()
	a := strings.Fields(scanner.Text())
	setA := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		setA[i], _ = strconv.Atoi(a[i])
	}

	scanner.Scan()
	b := strings.Fields(scanner.Text())
	setB := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		setB[i], _ = strconv.Atoi(b[i])
	}
	setBIndex := 0
	for i := 0; i < len(setA); i++ {
		for setBIndex < n && setB[setBIndex] < setA[i] {
			setBIndex++
		}
		if setBIndex < n && setB[setBIndex]-setA[i] <= r {
			fmt.Println(strconv.Itoa(setA[i]) + " " + strconv.Itoa(setB[setBIndex]))
		}
	}

}
