package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 回复数组序列
/**
题目描述
对于一个连续正整数组成的序列，可以将其拼接成一个字符串，再将字符串里的部分字符打乱顺序。如序列8 9 10 11 12，拼接成的字符串为89101112，打乱一部分字符后得到90811211，原来的正整数10就被拆成了0和1。
现给定一个按如上规则得到的打乱字符的字符串，请将其还原成连续正整数序列，并输出序列中最小的数字。

输入描述
输入一行，为打乱字符的字符串和正整数序列的长度，两者间用空格分隔，字符串长度不超过200，正整数不超过1000，保证输入可以还原成唯一序列。

输出描述
输出一个数字，为序列中最小的数字。
*/

/*
*输入：
19801211 5
输出：
8
*/
func main7() {
	//思路：将字符串s所有的字符存入map种，然后从1到1000-k+1开始每个都尝试一遍，将i~i+1的都转成字符串，
	//然后遍历字符串的字符串全部存入另外map种，最后判定第一个map和第二个map是否相同吗，相同则代表i即为序列中的最小值
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	s := arr[0]
	k, _ := strconv.Atoi(arr[1])
	mp := make(map[rune]int)
	for _, ch := range s {
		mp[ch] = mp[ch] + 1
	}
	ans := 0
	//从1000-k+1开始滑动
	for i := 1; i <= 1000-k+1; i++ {
		//滑动窗口，长度为k
		m := make(map[rune]int)
		for j := 0; j < k; j++ {
			rs := i + j
			//将rs转化成字符串
			rsStr := strconv.Itoa(rs)
			for _, ss := range rsStr {
				m[ss] = m[ss] + 1
			}
		}
		match := true
		for kk, v := range mp {
			if vv, exists := m[kk]; !exists || vv != v {
				match = false
				break
			}
		}
		if match {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
