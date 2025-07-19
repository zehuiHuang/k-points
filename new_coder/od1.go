package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main1() {
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//input := scanner.Text()

	input := "2.3,3,5.6,7,6;11,3,8.6,25,1;0.3,9,5.3,66,7.8;1,3,2,7,5;340,670,80.6;<=,<=,<="
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
