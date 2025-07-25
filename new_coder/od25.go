package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 斗地主之顺子
/**
题目描述
在斗地主扑克牌游戏中， 扑克牌由小到大的顺序为：3,4,5,6,7,8,9,10,J,Q,K,A,2，玩家可以出的扑克牌阵型有：单张、对子、顺子、飞机、炸弹等。

其中顺子的出牌规则为：由至少5张由小到大连续递增的扑克牌组成，且不能包含2。

例如：{3,4,5,6,7}、{3,4,5,6,7,8,9,10,J,Q,K,A}都是有效的顺子；而{J,Q,K,A,2}、 {2,3,4,5,6}、{3,4,5,6}、{3,4,5,6,8}等都不是顺子。

给定一个包含13张牌的数组，如果有满足出牌规则的顺子，请输出顺子。

如果存在多个顺子，请每行输出一个顺子，且需要按顺子的第一张牌的大小（必须从小到大）依次输出。

如果没有满足出牌规则的顺子，请输出No。

输入描述
13张任意顺序的扑克牌，每张扑克牌数字用空格隔开，每张扑克牌的数字都是合法的，并且不包括大小王：

2 9 J 2 3 4 K A 7 9 A 5 6

不需要考虑输入为异常字符的情况

输出描述
组成的顺子，每张扑克牌数字用空格隔开：

3 4 5 6 7
*/
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
	//思路：先对牌进行从小到大的排序
	//1、对牌进行分组，将最小的放入到数组中的groupA
	//2、从1～nums进行遍历，同时对分组进行遍历，找到一个分组时期符合条件：如果当前分组中最大的比当前值大1则表示将该值放入到当前分组，若每个都不满足该条件，则需要重新创建一个分组
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
				break
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
