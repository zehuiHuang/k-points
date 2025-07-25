package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 最长的指定瑕疵度的元音子串
/**
开头和结尾都是元音字母（aeiouAEIOU）的字符串为元音字符串，其中混杂的非元音字母数量为其瑕疵度。比如:

“a” 、 “aa”是元音字符串，其瑕疵度都为0
“aiur”不是元音字符串（结尾不是元音字符）
 “abira”是元音字符串，其瑕疵度为2
给定一个字符串，请找出指定瑕疵度的最长元音字符子串，并输出其长度，如果找不到满足条件的元音字符子串，输出0。

子串：字符串中任意个连续的字符组成的子序列称为该字符串的子串
*/

/*
*
输入：
0
asdbuiodevauufgh
输出：
3
*/
func main23() {
	//思路：滑动窗口
	//1、根据原字符串计算出元音字符的下标并组装成元音字符串（值为原字符串下标），滑动窗口双指针在元音字符串滑动
	//2、当双指针之间的瑕疵度小于n，则r++，若大于n，则l++，若等于n，则记录下来和之前历史的对比后选最大的
	//3、循环条件r<len(原音字符串)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := scanner.Text()
	ss := "aeiouAEIOU"
	mp := make(map[rune]bool)
	for i := 0; i < len(ss); i++ {
		mp[rune(ss[i])] = true
	}

	idx := []int{}
	//统计元音字符在字符串的下标并得出一个切片，在该切片上进行滑动
	for i := 0; i < len(s); i++ {
		if mp[rune(s[i])] {
			idx = append(idx, i)
		}
	}
	ans := 0

	l, r := 0, 0
	for r < len(idx) {
		//统计暇疵的数量:
		//计算思路：原数组的长 -元音字符串的长
		diff := (idx[r] - idx[l] + 1) - (r - l + 1)
		if diff == n {
			ans = max2(ans, idx[r]-idx[l]+1)
			r++
		} else if diff > n {
			l++
		} else {
			r++
		}
	}
	fmt.Println(ans)
}
