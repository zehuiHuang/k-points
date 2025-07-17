package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main10() {
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
