package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

//英文输入法
/**
题目描述
主管期望你来实现英文输入法单词联想功能。
需求如下：

依据用户输入的单词前缀，从已输入的英文语句中联想出用户想输入的单词，按字典序输出联想到的单词序列，
如果联想不到，请输出用户输入的单词前缀。
注意：

英文单词联想时，区分大小写
缩略形式如”don’t”，判定为两个单词，”don”和”t”
输出的单词序列，不能有重复单词，且只能是英文单词，不能有标点符号
输入描述
输入为两行。

首行输入一段由英文单词word和标点符号组成的语句str；

接下来一行为一个英文单词前缀pre。

0 < word.length() <= 20
0 < str.length <= 10000
0 < pre <= 20
输出描述
输出符合要求的单词序列或单词前缀，存在多个时，单词之间以单个空格分割
*/

/*
*输入：
I love you
He
输出：
He

输入
The furthest distance in the world, Is not between life and death, But when I stand in front of you, Yet you don’t know that I love you.
f
输出：
front furthest
*/
func main10() {
	//思路：1、对字符串处理（替换所有非字符串、去重、字典排序），遍历处理后的数据进行前缀匹配查询，用到strings.hasPrefix(原字符串，前缀字符串)
	scanner := bufio.NewScanner(os.Stdin)

	// 读取输入句子
	scanner.Scan()
	sentence := scanner.Text()

	// 读取前缀
	scanner.Scan()
	prefix := scanner.Text()

	// 替换所有非字母字符为空格
	reg := regexp.MustCompile(`[^a-zA-Z]`)
	sentence = reg.ReplaceAllString(sentence, " ")

	// 分割单词并去重
	wordSet := make(map[string]bool)
	words := strings.Fields(sentence)
	for _, word := range words {
		wordSet[word] = true
	}

	// 获取去重后的单词并排序
	uniqueWords := make([]string, 0, len(wordSet))
	for word := range wordSet {
		uniqueWords = append(uniqueWords, word)
	}
	sort.Strings(uniqueWords)

	// 收集匹配前缀的单词
	var result []string
	for _, word := range uniqueWords {
		if strings.HasPrefix(word, prefix) {
			result = append(result, word)
		}
	}

	// 输出结果
	if len(result) > 0 {
		fmt.Println(strings.Join(result, " "))
	} else {
		fmt.Println(prefix)
	}
}
