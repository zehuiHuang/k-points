/**
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，
并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；
反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换
*/
/**
输入：
3
12abc-abCABc-4aB@
输出：
12abc-abc-ABC-4aB-@


12
12abc-abCABc-4aB@
12abc-abCABc4aB@
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// 字符串分隔
/**
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换

输入描述
输入为两行，第一行为参数K，第二行为字符串S

输出描述
输出转换后的字符串
*/

/*
*
输入：
3
12abc-abCABc-4aB@
输出：
12abc-abc-ABC-4aB-@

输入：
12
12abc-abCABc-4aB@
输出：
12abc-abCABc4aB@
*/
func main20() {
	//思路：1、将字符串s按照-进行切分生成切片，则切片下标为0的字符串可以直接放入结果里
	//2、将剩下的字符串进行循环，步长为k，对长度为k的字符串进行处理
	//3、分别统计长度为k的字符串小写和大写的数量，若小写多则将大写转小写，若大写多则小写转大写，若相同则不处理

	//这里用到了一些字符串传的基本方法：
	//unicode.IsLower()
	//unicode.IsUpper()
	//unicode.ToLower()
	//unicode.ToUpper()
	//strings.Builder的WriteRune

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	S := scanner.Text() // 读取字符串S

	// 分割字符串
	parts := strings.Split(S, "-")
	if len(parts) == 0 {
		fmt.Println(S)
		return
	}

	// 第一个子串保持不变
	firstPart := parts[0]
	// 合并剩余子串
	merged := strings.Join(parts[1:], "")

	// 处理合并后的字符串
	var newParts []string
	newParts = append(newParts, firstPart)

	// 按K长度分组
	for i := 0; i < len(merged); i += K {
		end := i + K
		if end > len(merged) {
			end = len(merged)
		}
		group := merged[i:end]
		newParts = append(newParts, processGroup(group))
	}

	// 构建结果字符串
	result := strings.Join(newParts, "-")
	fmt.Println(result)
}

func processGroup(s string) string {
	lowerCount := 0
	upperCount := 0

	// 统计大小写字母数量
	for _, ch := range s {
		if unicode.IsLower(ch) {
			lowerCount++
		} else if unicode.IsUpper(ch) {
			upperCount++
		}
	}

	// 根据统计结果转换
	var builder strings.Builder
	if lowerCount > upperCount {
		for _, ch := range s {
			if unicode.IsUpper(ch) {
				builder.WriteRune(unicode.ToLower(ch))
			} else {
				builder.WriteRune(ch)
			}
		}
	} else if upperCount > lowerCount {
		for _, ch := range s {
			if unicode.IsLower(ch) {
				builder.WriteRune(unicode.ToUpper(ch))
			} else {
				builder.WriteRune(ch)
			}
		}
	} else {
		builder.WriteString(s)
	}

	return builder.String()
}
