package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main6() {
	//凑成二位数组
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrP := strings.Fields(scanner.Text())

	r, _ := strconv.Atoi(arrP[0])
	c, _ := strconv.Atoi(arrP[1])
	//s, _ := strconv.Atoi(arrP[2])
	//m, _ := strconv.Atoi(arrP[3])

	matrix := make([][]int, r)

	for i := 0; i < r; i++ {
		scanner.Scan()
		arr := strings.Fields(scanner.Text())
		matrix[i] = make([]int, c)
		for j := 0; j < c; j++ {
			v, _ := strconv.Atoi(arr[j])
			matrix[i][j] = v
		}
	}

}
