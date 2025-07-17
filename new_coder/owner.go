package main

import (
	"bufio"
	"fmt"
	"os"
)

func main1111111() {
	//接受输入的字符串
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()
	input2 := scanner.Text()
	fmt.Println(input1, input2)

}
