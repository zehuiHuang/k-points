package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//字符串切分

/*
*
题目描述：构成指定长度字符串的个数 (本题分值100)
给定 M（0 < M ≤ 30）个字符（a-z），从中取出任意字符（每个字符只能用一次）拼接成长度为 N（0 < N ≤ 5）的字符串，

要求相同的字符不能相邻，计算出给定的字符列表能拼接出多少种满足条件的字符串，

输入非法或者无法拼接出满足条件的字符串则返回0。
*/

/*
*
输入描述
给定的字符列表和结果字符串长度，中间使用空格(" ")拼接

输出描述
满足条件的字符串个数
*/

/*
*用例1
输入：
aab 2
输出
2
*/

// 思路：使用回溯方法
func main33() {
	//思路：使用回溯算饭，遍历s中的每个字符串，由于是排序问题，所有没用index，而是每次循环都是从0～len(nums)中选择
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	s := input[0]
	k, _ := strconv.Atoi(input[1])

	//存储每个字符串的数量（也有去重的作用）
	mp := make(map[string]int)
	used := make([]bool, len(s))

	var dfs func(s, current string, mp map[string]int, c []bool)
	//s:指定的字符串；mp：存储符合条件的字符串；used:每次从指定字符串取值时判定是否已经被取过
	dfs = func(s, current string, mp map[string]int, used []bool) {
		if len(current) == k {
			mp[current] = 1
			return
		}
		for i := 0; i < len(s); i++ {
			c := s[i]
			//如果字符串已经被用过，或者数组i对应的字符串 和current字符串的末尾的字符串相同，则跳过该层遍历
			if used[i] || (len(current) > 0 && c == current[len(current)-1]) {
				continue
			}
			used[i] = true
			dfs(s, current+string(c), mp, used)
			used[i] = false
		}
	}
	dfs(s, "", mp, used)
	fmt.Println(len(mp))

}
