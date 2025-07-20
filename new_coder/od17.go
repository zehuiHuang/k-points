package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
*题目描述
给定一个整型数组，请从该数组中选择3个元素组成最小数字并输出

（如果数组长度小于3，则选择数组中所有元素来组成最小数字）。

输入描述
一行用半角逗号分割的字符串记录的整型数组，0 < 数组长度 <= 100，0 < 整数的取值范围 <= 10000。

输出描述
由3个元素组成的最小数字，如果数组长度小于3，则选择数组中所有元素来组成最小数字。

输入：
21,30,62,5,31
输出：
21305
*/
func main17() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Split(scanner.Text(), ",")
	//arr := make([]int, len(input))
	//for i := range input {
	//	arr[i], _ = strconv.Atoi(input[i])
	//}

	sort.Slice(arr, func(i, j int) bool {
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[j])
		return a < b
	})
	if len(arr) > 3 {
		arr = arr[:3]
	}
	//给三个数进行组合
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]+arr[j] < arr[j]+arr[i]
	})
	res := ""
	for i := range arr {
		res = res + arr[i]
	}
	fmt.Println(res)
}
