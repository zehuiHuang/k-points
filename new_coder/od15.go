package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
*

输入	/acm,/bb
输出	/acm/bb
*/
func main15() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputArr := strings.Split(scanner.Text(), ",")
	start := "/"
	end := "/"
	if len(inputArr) > 0 && len(inputArr[0]) > 0 {
		start = inputArr[0]
	}
	if len(inputArr) > 1 && len(inputArr[1]) > 0 {
		end = inputArr[1]
	}
	start = strings.TrimRight(start, "/")
	end = strings.TrimLeft(end, "/")
	fmt.Println(start + "/" + end)
}
