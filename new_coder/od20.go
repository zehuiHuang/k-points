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
	"strings"
	"unicode"
)

func main20() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var K int
	fmt.Sscanf(scanner.Text(), "%d", &K) // 读取K值
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
