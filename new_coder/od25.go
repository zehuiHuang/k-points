package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 斗地主之顺子

/*
*
输入：
2 9 J 2 3 4 K A 7 9 A 5 6
输出
3 4 5 6 7

输入：
2 9 J 10 3 4 K A 7 Q A 5 6
输出：
3 4 5 6 7
9 10 J Q K A
*/
func main25() {
	mp := make(map[string]int)
	mp["3"] = 3
	mp["4"] = 4
	mp["5"] = 5
	mp["6"] = 6
	mp["7"] = 7
	mp["8"] = 8
	mp["9"] = 9
	mp["10"] = 10
	mp["J"] = 11
	mp["Q"] = 12
	mp["K"] = 13
	mp["A"] = 14
	mp["2"] = 16
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Fields(scanner.Text())
	//思路：先按照大小排序，然后把第一个放入一个数组中，然后按照顺序（相差1）进行放入，若相差不等于1，则重新开一个数组继续这个逻辑
	sort.Slice(arr, func(i, j int) bool {
		return mp[arr[i]] < mp[arr[j]]
	})
	rets := [][]string{}
	ret := []string{arr[0]}
	rets = append(rets, ret)

	for i := 1; i < len(arr); i++ {
		inq := false
		for j := 0; j < len(rets); j++ {
			q := rets[j]
			if len(q) > 0 && mp[arr[i]]-mp[q[len(q)-1]] == 1 {
				rets[j] = append(rets[j], arr[i])
				inq = true
			}
		}
		if !inq {
			temp := []string{arr[i]}
			rets = append(rets, temp)
		}
	}
	sort.Slice(rets, func(i, j int) bool {
		return rets[i][0] < rets[j][0]
	})
	if len(rets) == 0 {
		fmt.Println("No")
	}
	//找到个数大于等于5的顺子
	for i := 0; i < len(rets); i++ {
		if len(rets[i]) >= 5 {
			fmt.Println(strings.Join(rets[i], " "))
		}
	}
}
