package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main29() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取源字符串
	scanner.Scan()
	source := scanner.Text()

	// 读取目标字符串
	scanner.Scan()
	target := scanner.Text()

	// 处理可选段标记（将 [...] 转换为正则表达式字符组）
	targetRegex := regexp.MustCompile(`\[(.*?)\]`)
	target = targetRegex.ReplaceAllString(target, "[$1]")

	// 编译正则表达式
	re, err := regexp.Compile(target)
	if err != nil {
		fmt.Println(-1)
		return
	}

	// 查找匹配位置
	loc := re.FindStringIndex(source)
	if loc == nil {
		fmt.Println(-1)
	} else {
		fmt.Println(loc[0])
	}
}
