package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// 主管期望你来实现英文输入法单词联想功能。
func main() {
	//The furthest distance in the world, Is not between life and death, But when I stand in front of you, Yet you don’t know that I love you.
	//f
	//front furthest

	//获取输入信息
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := scanner.Text()

	scanner.Scan()
	input2 := scanner.Text()

	reg := regexp.MustCompile(`[^a-zA-Z]`)
	sentence := reg.ReplaceAllString(input1, " ")

	words := strings.Fields(sentence)
	//去重
	mp := make(map[string]bool)
	//
	arr := []string{}
	for _, s := range words {
		mp[s] = true
	}
	for v := range mp {
		arr = append(arr, v)
	}
	sort.Strings(arr)

	ans := []string{}
	for _, v := range arr {
		if strings.HasPrefix(v, input2) {
			ans = append(ans, v)
		}
	}

	if len(ans) > 0 {
		fmt.Println(strings.Join(ans, " "))
	} else {
		fmt.Println(input2)
	}

}
