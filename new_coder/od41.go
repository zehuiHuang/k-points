package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

//书籍叠放

/*
*
题目描述
书籍的长、宽都是整数对应 (l,w)。如果书A的长宽度都比B长宽大时，则允许将B排列放在A上面。现在有一组规格的书籍，书籍叠放时要求书籍不能做旋转，请计算最多能有多少个规格书籍能叠放在一起。

输入描述
输入：books = [[20,16],[15,11],[10,10],[9,10]]

说明：总共4本书籍，第一本长度为20宽度为16；第二本书长度为15宽度为11，依次类推，最后一本书长度为9宽度为10.

输出描述
输出：3

说明: 最多3个规格的书籍可以叠放到一起, 从下到上依次为: [20,16],[15,11],[10,10]
*/

/*
*输入
[[20,16],[15,11],[10,10],[9,10]]
输出
3
*/
/**
思路：先按照某一纬度进行生序A排列，然后在另外一个纬度进行 最长递增子序列的查询
最长递增子序列的查询的方案为动态规划：定义以i为结尾的最长子序列的长度,那么就是查询从0～i-1之间所有的max(dp[j]+1)即为dp[i] ，
但是要求是v[i]>v[j]
*/
func main41() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// 原始字符串数据
	strData := scanner.Text()

	// 创建二维切片接收解析结果
	var tables [][]int

	// 使用 json.Unmarshal 解析 JSON 格式的字符串
	err := json.Unmarshal([]byte(strData), &tables)
	if err != nil {
		return
	}
	//以长度进行升序排序
	sort.Slice(tables, func(i, j int) bool {
		return tables[i][0] < tables[j][0]
	})

	//最长递增子序列的查询
	//定义高度h，以i为结尾的最长子序列的长度为dp[i]
	dp := make([]int, len(tables))
	//初始化
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(tables); i++ {
		for j := 0; j < i; j++ {
			if tables[i][1] > tables[j][1] {
				dp[i] = max2(dp[i], dp[j]+1)
			}
		}
	}
	index := len(tables) - 1
	fmt.Println(dp[index])
}
