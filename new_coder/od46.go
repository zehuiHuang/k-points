package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//最长的顺子
/**
题目描述
斗地主起源于湖北十堰房县，据说是一位叫吴修全的年轻人根据当地流行的扑克玩法“跑得快”改编的，如今已风靡整个中国，并流行于互联网上。
牌型：单顺，又称顺子，最少5张牌，最多12张牌(3…A)不能有2，也不能有大小王，不计花色。
例如： 3-4-5-6-7-8，7-8-9-10-J-Q，3-4-5-6-7-8-9-10-J-Q-K-A
可用的牌 3<4<5<6<7<8<9<10<J<Q<K<A<2<B(小王)<C(大王)，每种牌除大小王外有四种花色
(共有13×4+2张牌)

输入：
手上有的牌
已经出过的牌(包括对手出的和自己出的牌)

输出：
对手可能构成的最长的顺子(如果有相同长度的顺子，输出牌面最大的那一个)，
如果无法构成顺子，则输出 NO-CHAIN。

输入描述
输入的第一行为当前手中的牌
输入的第二行为已经出过的牌

输出描述
最长的顺子
*/

/*
*输入
3-3-3-3-4-4-5-5-6-7-8-9-10-J-Q-K-A
4-5-6-7-8-8-8
输出
9-10-J-Q-K-A
*/
func main() {
	//思路：
	//1、算出对方的牌：总牌-我的牌-打出的牌
	//2、基于对方的牌，查询出最长最大的顺子
	//创建一个切片表示牌，索引表示牌面数，值为牌的个数
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	myCards := strings.Split(scanner.Text(), "-")

	scanner.Scan()
	usedCards := strings.Split(scanner.Text(), "-")

	//3-4-5-6-7-8-9-10-J-Q-K-A 2-B-C
	count := []int{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 0, 4, 1, 1}

	mp1 := map[int]string{
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "10",
		11: "J",
		12: "Q",
		13: "K",
		14: "A",
		16: "2",
		17: "B",
		18: "C",
	}

	mp2 := map[string]int{
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"J":  11,
		"Q":  12,
		"K":  13,
		"A":  14,
		"2":  16,
		"B":  17,
		"C":  18,
	}
	//总牌减去我的牌
	for i := range myCards {
		count[mp2[myCards[i]]] -= 1
	}

	//总牌减去已出的牌
	for i := range usedCards {
		count[mp2[usedCards[i]]] -= 1
	}

	//从左边开始能组成5张的 返回一定在3～10之间（左右都是闭区间）
	l := 3
	r := 10
	maxLen := 0
	ans := ""
	for l < r {
		temp := []string{}
		for i := l; i < 16; i++ {
			if count[i] > 0 {
				temp = append(temp, mp1[i])
			} else {
				if len(temp) >= 5 && len(temp) >= maxLen {
					maxLen = len(temp)
					ans = strings.Join(temp, "-")
				}
				l = i
				break
			}
		}
		l++
	}
	fmt.Println(ans)
}
