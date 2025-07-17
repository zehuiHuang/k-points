package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	//2.3,3,5.6,7,6;11,3,8.6,25,1;0.3,9,5.3,66,7.8;1,3,2,7,5;340,670,80.6;<=,<=,<=
	//打印 false 458
	arr := strings.Split(input, ";")
	if len(arr) < 6 {
		fmt.Println("false 0")
		return
	}
	//二位数组
	arr2 := make([][]string, len(arr))
	for i, v := range arr {
		arr2[i] = strings.Split(v, ",")
	}
	//具体的值
	a0 := convertFload(arr2[0])
	a1 := convertFload(arr2[1])
	a2 := convertFload(arr2[2])
	//乘的值
	a3 := convertFload(arr2[3])
	//不等式的右侧结果
	a4 := convertFload(arr2[4])
	//不等式符号
	a5 := arr2[5]
	//转化为fload65
	b1 := a0[0]*a3[0] + a0[1]*a3[1] + a0[2]*a3[2] + a0[3]*a3[3] + a0[4]*a3[4] - a4[0]
	b2 := a1[0]*a3[0] + a1[1]*a3[1] + a1[2]*a3[2] + a1[3]*a3[3] + a1[4]*a3[4] - a4[1]
	b3 := a2[0]*a3[0] + a2[1]*a3[1] + a2[2]*a3[2] + a2[3]*a3[3] + a2[4]*a3[4] - a4[2]
	flag := condition(b1, a5[0]) && condition(b2, a5[1]) && condition(b3, a5[2])
	v := max(b1, b2, b3)
	fmt.Printf("%t,%v\n", flag, int(v))
}

func convertFload(arr []string) []float64 {
	ans := make([]float64, len(arr))
	for i, s := range arr {
		res, _ := strconv.ParseFloat(s, 64)
		ans[i] = res
	}
	return ans
}

func condition(b float64, p string) bool {
	switch p {
	case ">":
		return b > 0
	case ">=":
		return b >= 0
	case "=":
		return b == 0
	case "<":
		return b < 0
	case "<=":
		return b <= 0
	default:
		return false
	}
}

func max(a, b, c float64) float64 {
	ans := 0.0
	if a > b {
		ans = a
	} else {
		ans = b
	}
	if ans < c {
		ans = c
	}
	return ans
}
