package main

//字符串序列判定
import (
	"bufio"
	"fmt"
	"os"
)

/**
输入两个字符串 S 和 L ，都只包含英文小写字母。S长度 ≤ 100，L长度 ≤ 500,000。判定S是否是L的有效子串。

判定规则：S 中的每个字符在 L 中都能找到（可以不连续），且 S 在Ｌ中字符的前后顺序与 S 中顺序要保持一致。（例如，S = ”ace” 是 L= ”abcde” 的一个子序列且有效字符是a、c、e，而”aec”不是有效子序列，且有效字符只有a、e）

输入描述：
输入两个字符串 S 和 L，都只包含英文小写字母。S长度 ≤ 100，L长度 ≤ 500,000。

先输入S，再输入L，每个字符串占一行。

输出描述：
S 串最后一个有效字符在 L 中的位置。（首位从0开始计算，无有效字符返回-1）
*/

/*
*
输入：
ace
abcde
输出：
4
*/
func main21() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	l := scanner.Text()
	//思路：滑动窗口
	i, j := 0, 0
	for i < len(s) && j < len(l) {
		if s[i] == l[j] {
			i++
			j++
		} else {
			j++
		}
	}
	if i == len(s) {
		fmt.Println(j - 1)
	} else {
		fmt.Println(-1)
	}
}
