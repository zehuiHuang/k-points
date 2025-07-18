package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*
*题目描述
给出3组点坐标(x, y, w, h)，-1000<x,y<1000，w,h为正整数。

(x, y, w, h)表示平面直角坐标系中的一个矩形：

x, y为矩形左上角坐标点，w, h向右w，向下h。

(x, y, w, h)表示x轴(x, x+w)和y轴(y, y-h)围成的矩形区域；

(0, 0, 2, 2)表示 x轴(0, 2)和y 轴(0, -2)围成的矩形区域；

(3, 5, 4, 6)表示x轴(3, 7)和y轴(5, -1)围成的矩形区域；

求3组坐标构成的矩形区域重合部分的面积。

输入描述
3行输入分别为3个矩形的位置，分别代表“左上角x坐标”，“左上角y坐标”，“矩形宽”，“矩形高” -1000 <= x,y < 1000
*/
func main8() {
	/**
	  	输入
	    1 6 4 4
	    3 5 3 4
	    0 3 7 3
	    输出
	    2
	*/

	x_coords := []int{}
	y_coords := []int{}
	rectangles := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 3; i++ {
		scanner.Scan()
		a := strings.Fields(scanner.Text())
		x1, _ := strconv.Atoi(a[0])
		y1, _ := strconv.Atoi(a[1])
		w, _ := strconv.Atoi(a[2])
		h, _ := strconv.Atoi(a[3])

		x2 := x1 + w
		y2 := y1 - h
		x_coords = append(x_coords, x1)
		x_coords = append(x_coords, x2)

		y_coords = append(y_coords, y1)
		y_coords = append(y_coords, y2)
		rectangles = append(rectangles, []int{x1, x2, y1, y2})
	}

}
