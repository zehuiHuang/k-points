package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//判断一组不等式是否满足约束并输出最大差

/*
*
题目描述
给定一组不等式，判断是否成立并输出不等式的最大差(输出浮点数的整数部分)

要求:

不等式系数为 double类型，是一个二维数组
不等式的变量为 int类型，是一维数组;
不等式的目标值为 double类型，是一维数组
不等式约束为字符串数组，只能是:“>”,“>=”,“<”,“<=”,“=”，
例如，不等式组:

a11x1+a12x2+a13x3+a14x4+a15x5<=b1;

a21x1+a22x2+a23x3+a24x4+a25x5<=b2;

a31x1+a32x2+a33x3+a34x4+a35x5<=b3;

最大差 = max{(a11x1+a12x2+a13x3+a14x4+a15x5-b1),(a21x1+a22x2+a23x3+a24x4+ a25x5-b2),(a31x1+a32x2+a33x3+a34x4+a35x5-b3)},

类型为整数(输出浮点数的整数部分)

输入描述
a11,a12,a13,a14,a15,a21,a22,a23,a24,a25, a31,a32,a33,a34,a35,x1,x2,x3,x4,x5,b1,b2,b3,<=,<=,<=

1)不等式组系数(double类型):

a11,a12,a13,a14,a15

a21,a22,a23,a24,a25

a31,a32,a33,a34,a35

2)不等式变量(int类型):x1,x2,x3,x4,x5

3)不等式目标值(double类型):b1,b2,b3

4)不等式约束(字符串类型):<=,<=,<=

输出描述
true或者 false，最大差
*/

/*
*
输入
2.3,3,5.6,7,6;11,3,8.6,25,1;0.3,9,5.3,66,7.8;1,3,2,7,5;340,670,80.6;<=,<=,<=
输出
false 458
*/
func main1() {
	//思路：按照题目提示步骤操作即可（注意float类型转化）
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	// 分割输入字符串
	parts := strings.Split(input, ";")
	if len(parts) < 6 {
		fmt.Println("false 0")
		return
	}

	// 解析二维数组
	arr := make([][]string, len(parts))
	for i, p := range parts {
		arr[i] = strings.Split(p, ",")
	}

	// 解析各个数组
	a1 := parseDoubleArray(arr[0])
	a2 := parseDoubleArray(arr[1])
	a3 := parseDoubleArray(arr[2])
	x := parseDoubleArray(arr[3])
	b := parseDoubleArray(arr[4])
	y := arr[5] // 约束条件数组

	// 计算差值
	diff1 := a1[0]*x[0] + a1[1]*x[1] + a1[2]*x[2] + a1[3]*x[3] + a1[4]*x[4] - b[0]
	diff2 := a2[0]*x[0] + a2[1]*x[1] + a2[2]*x[2] + a2[3]*x[3] + a2[4]*x[4] - b[1]
	diff3 := a3[0]*x[0] + a3[1]*x[1] + a3[2]*x[2] + a3[3]*x[3] + a3[4]*x[4] - b[2]

	// 检查约束条件
	flag := compareWithZero(diff1, y[0]) &&
		compareWithZero(diff2, y[1]) &&
		compareWithZero(diff3, y[2])

	// 找出最大差值
	maxDiff := max(diff1, diff2, diff3)

	// 输出结果
	fmt.Printf("%t %d\n", flag, int(maxDiff))
}

// 将字符串数组转换为float64数组
func parseDoubleArray(strArr []string) []float64 {
	arr := make([]float64, len(strArr))
	for i, s := range strArr {
		arr[i], _ = strconv.ParseFloat(s, 64)
	}
	return arr
}

// 根据约束条件比较值与0
func compareWithZero(val float64, constraint string) bool {
	switch constraint {
	case ">":
		return val > 0
	case ">=":
		return val >= 0
	case "<":
		return val < 0
	case "<=":
		return val <= 0
	case "=":
		return val == 0
	default:
		return false
	}
}

// 找出三个float64中的最大值
func max(a, b, c float64) float64 {
	maxVal := a
	if b > maxVal {
		maxVal = b
	}
	if c > maxVal {
		maxVal = c
	}
	return maxVal
}
