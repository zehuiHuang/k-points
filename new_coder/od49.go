package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//荒岛求生
/**
题目描述
一个荒岛上有若干人，岛上只有一条路通往岛屿两端的港口，大家需要逃往两端的港口才可逃生。

假定每个人移动的速度一样，且只可选择向左或向右逃生。

若两个人相遇，则进行决斗，战斗力强的能够活下来，并损失掉与对方相同的战斗力；若战斗力相同，则两人同归于尽。

输入描述
给定一行非 0 整数数组，元素个数不超过30000；

正负表示逃生方向（正表示向右逃生，负表示向左逃生），绝对值表示战斗力，越左边的数字表示里左边港口越近，逃生方向相同的人永远不会发生决斗。

输出描述
能够逃生的人总数，没有人逃生输出0，输入异常时输出-1。
*/

/*
*输入
5 10 8 -8 -5

输出
2
*/
func main49() {
	//思路：遍历当前数组，如果当前值小于零，那么和栈中的栈顶元素对比，如果相同，则同归于尽（当前元素忽略+出栈），如果栈顶元素大，
	//则当前值忽略，同时栈顶元素需要减去当前值，如果当前元素大，则当前元素减去栈顶元素的值，并且栈顶元素出栈

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())

	tables := make([]int, len(input))
	for i := range tables {
		tables[i], _ = strconv.Atoi(input[i])
	}
	stack := []int{}
	for i := range tables {
		alive := true
		for alive && tables[i] < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {

			if stack[len(stack)-1] == -tables[i] {
				//当前人被同归于尽
				alive = false
				//出栈，栈顶人被同归于尽
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] > -tables[i] {
				//当前人被杀掉
				alive = false
				//栈顶人被消耗
				stack[len(stack)-1] = stack[len(stack)-1] + tables[i]
			} else {
				//当前人存活，但是被栈顶人消耗了
				tables[i] = tables[i] + stack[len(stack)-1]
				//栈顶人被杀掉
				stack = stack[:len(stack)-1]
			}

		}
		if alive {
			stack = append(stack, tables[i])
		}
	}
	fmt.Println(len(stack))
}
