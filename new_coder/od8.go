package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//矩形相交面积

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
	//思路：
	//1、找出能囊括所有矩形的最小二位数组（m*n）：
	//计算方式是找到所有点的x坐标和y坐标， 其中最大的x坐标+offset即为二维数组的行数， 最大y+offset为数组的列数 （为了能从数组下标0开始，设置了偏移量offset）
	//2、找到每个矩阵对应的二维数组，行数：从min(x1,x2)+offset 到 max(x1,x2)+offset,列数：从min(y1,y2)+offset到max(y1,y2)+offset
	//3、遍历每个矩形矩阵，将值加+1
	//4、遍历囊括所有矩阵的位置，值为3的即为三个矩阵都覆盖的点
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
	//找到最大矩阵
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for i := range x_coords {
		minX = min(minX, x_coords[i])
		maxX = max2(maxX, x_coords[i])
	}
	for i := range y_coords {
		minY = min(minY, y_coords[i])
		maxY = max2(maxY, y_coords[i])
	}
	//计算偏移量：为了能让二位数组下标从0开始
	offsetX := 0 - minX
	offsetY := 0 - minY

	tables := make([][]int, int(math.Abs(float64(maxX-minX))))
	for i := range tables {
		tables[i] = make([]int, int(math.Abs(float64(maxY-minY))))
	}
	for i := range rectangles {
		x1 := rectangles[i][0]
		x2 := rectangles[i][1]
		y1 := rectangles[i][2]
		y2 := rectangles[i][3]
		minX := min(x1, x2)
		maxX := max2(x1, x2)

		minY := min(y1, y2)
		maxY := max2(y1, y2)
		//遍历每个矩形的位置并+1
		for j := minX + offsetX; j < maxX+offsetX; j++ {
			for k := minY + offsetY; k < maxY+offsetY; k++ {
				tables[j][k] += 1
			}
		}
	}
	ans := 0
	for i := range tables {
		for j := range tables[i] {
			if tables[i][j] == 3 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
