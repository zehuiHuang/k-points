package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//区间集合
/**
题目描述
给定一组闭区间，其中部分区间存在交集。

任意两个给定区间的交集，称为公共区间(如:[1,2],[2,3]的公共区间为[2,2]，[3,5],[3,6]的公共区间为[3,5])。

公共区间之间若存在交集，则需要合并(如:[1,3],[3,5]区间存在交集[3,3]，需合并为[1,5])。

按升序排列输出合并后的区间列表。

输入描述
一组区间列表，

区间数为 N: 0<=N<=1000;

区间元素为 X: -10000<=X<=10000。

输出描述
升序排列的合并区间列表
*/

/*
*输入：
4
0 3
1 3
3 5
3 6
输出：
1 5

输入：
4
0 3
1 4
4 7
5 8
输出：
1 3
4 4
5 7
*/
func main39() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())

	table := make([][2]int, num)

	for i := 0; i < num; i++ {
		scanner.Scan()
		rows := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(rows[0])
		r, _ := strconv.Atoi(rows[1])
		table[i] = [2]int{l, r}
	}
	//升序排序
	sort.Slice(table, func(i, j int) bool {
		return table[i][0] < table[j][0]
	})
	ans := [][]int{}

	//找到公共区间
	for i := 0; i < num; i++ {
		_, r1 := table[i][0], table[i][1]
		for j := i + 1; j < num; j++ {
			l2, r2 := table[j][0], table[j][1]
			if l2 > r1 {
				continue
			} else { //l2<=r1
				ans = append(ans, []int{l2, min(r1, r2)})
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][1]
	})
	res := [][]int{ans[0]}
	//对公告区间进行合并
	for i := 1; i < len(ans); i++ {
		_, br := res[len(res)-1][0], res[len(res)-1][1]
		l, r := ans[i][0], ans[i][1]
		if l <= br {
			res[len(res)-1][1] = r
		} else {
			res = append(res, ans[i])
		}
	}

	for i := range res {
		fmt.Println(strconv.Itoa(res[i][0]) + " " + strconv.Itoa(res[i][1]))
	}
}
