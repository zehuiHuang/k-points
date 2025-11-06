package practice

import (
	"strings"
	"unicode"
)

// b3[a[2[c]]] => baccaccacc

//[3]

type Pair struct {
	//
	K int
	S string
}

// 思路:用栈模拟递归
// 1、如果是数字和字母,则记录数字,并累积字母
// 2、如果是‘[’,则将之前累积数字和字母封装成Pair对象 并入栈
// 3、如果是‘]’,则表示要出栈,并配合,数字和字母进行计算,拼装成字符串
func decodeString(s string) string {
	//模拟栈进行递归
	stack := []Pair{}
	res := ""
	k := 0
	for _, curr := range s {
		//数字
		if curr >= '0' && curr <= '9' {
			//防止数字是多位的,比如23[a]
			k = k*10 + int(curr-'0')
		} else if (curr >= 'a' && curr <= 'z') || (curr >= 'A' && curr <= 'Z') { //字母
			res += string(curr)
		} else if curr == '[' { //左括号
			//入栈
			pair := Pair{k, res}
			stack = append(stack, pair)
			//初始化k和res
			k = 0
			res = ""
		} else { //右括号
			//出栈并封装结果
			//3[a[2[c]]]  => 2,c => cc
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//关键: 要理解Pair结构体中的K和S字段的含义, K表示res的数量(即后面字符串累积的数量),S不需要乘以倍数,需要s+res*num
			res = v.S + strings.Repeat(res, v.K)
		}
	}
	return res
}

func decodeString2(s string) string {
	type pair struct {
		s string
		k int
	}
	stack := []pair{} // 用于模拟计算机的递归
	res := ""
	k := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			res += string(c)
		} else if unicode.IsDigit(c) {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			// 模拟递归
			// 在递归之前，把当前递归函数中的局部变量 res 和 k 保存到栈中
			stack = append(stack, pair{res, k})
			// 递归，初始化 res 和 k
			res = ""
			k = 0
		} else { // ']'
			// 递归结束，从栈中恢复递归之前保存的局部变量
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 此时 res 是下层递归的返回值，将其重复 p.k 次，拼接到递归前的 p.s 之后
			res = p.s + strings.Repeat(res, p.k)
		}
	}
	return res
}
