package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//拼接URL

/**
题目描述
给定一个url前缀和url后缀,通过,分割 需要将其连接为一个完整的url

如果前缀结尾和后缀开头都没有/，需要自动补上/连接符
如果前缀结尾和后缀开头都为/，需要自动去重
约束：不用考虑前后缀URL不合法情况

输入描述
url前缀(一个长度小于100的字符串)，url后缀(一个长度小于100的字符串)

输出描述
拼接后的ur
*/
/*
*
输入	/acm,/bb
输出	/acm/bb
*/
func main15() {
	//思路：将第一个去掉右边的"/",然后去掉第二个的左边的"/",最后left+"/"+right
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
