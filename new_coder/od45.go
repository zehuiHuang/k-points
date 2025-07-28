package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//最小循环子数组

/*
*题目描述
给定一个由若干整数组成的数组nums，请检查数组是否是由某个子数组重复循环拼接而成，请输出这个最小的子数组。

输入描述
第一行输入数组中元素个数n，1 ≤ n ≤ 100000

第二行输入数组的数字序列nums，以空格分割，0 ≤ nums[i] < 10

输出描述
输出最小的子数组的数字序列，以空格分割；

备注
数组本身是其最大的子数组，循环1次可生成的自身；
*/

/*
*输入:
9
1 2 1 1 2 1 1 2 1
输出:
1 2 1
说明:
数组[1,2,1,1,2,1,1,2,1] 可由子数组[1,2,1]重复循环3次拼接而成
*/
func main45() {
	//思路：1、求出字符串的最长相同前后缀的长度m
	//2、最小循环子数组为nums[:n-m]
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	nums := strings.Fields(scanner.Text())

	m := getNext(nums)

	if n%(n-m) == 0 {
		fmt.Println(strings.Join(nums[:n-m], " "))
	} else {
		fmt.Println(strings.Join(nums, " "))
	}
}

func getNext(nums []string) int {
	ans := 0
	l, r := 1, len(nums)-2
	for l < len(nums)-1 {
		left := nums[:l]
		right := nums[r+1:]
		if strings.Join(left, "") == strings.Join(right, "") {
			ans = max2(ans, len(right))
		}
		l++
		r--
	}
	return ans
}
