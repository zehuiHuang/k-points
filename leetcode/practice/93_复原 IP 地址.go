package practice

import "strconv"

/**
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。



示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

*/
// 93. 复原 IP 地址
// 思路：回溯算法，通过一个个决策树来切割
func restoreIpAddresses(s string) []string {
	res := []string{}
	var dfs func(s string, tpms []string)
	//决策树
	dfs = func(s string, tpms []string) {
		//判定是否满足条件
		if s == "" && len(tpms) != 4 {
			return
		}
		if s != "" && len(tpms) > 4 {
			return
		}
		//对符合条件的进行收集
		if s == "" && len(tpms) == 4 {
			tmp := tpms[0]
			for i := 1; i < len(tpms); i++ {
				tmp += "." + tpms[i]
			}
			res = append(res, tmp)
			return
		}
		//递归:对字符串进行切割,从第一个开始,到第三个结束(这是因为在切割就超过了3位,而IP某位置最长才3)
		for i := 1; i <= 3 && i <= len(s); i++ {
			subStr := s[:i]
			if len(subStr) == 0 || subStr[0] == '0' {
				break
			}
			num, _ := strconv.Atoi(subStr)
			if num > 255 {
				break
			}
			dfs(s[i:], append(tpms, subStr))
		}

	}
	tpms := []string{}
	dfs(s, tpms)
	return res
}
