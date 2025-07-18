package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main7() {
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
