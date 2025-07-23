package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//光伏场地建设规划
/*
*
题目描述
祖国西北部有一片大片荒地，其中零星的分布着一些湖泊，保护区，矿区;
整体上常年光照良好，但是也有一些地区光照不太好。

某电力公司希望在这里建设多个光伏电站，生产清洁能源对每平方公里的土地进行了发电评估，
其中不能建设的区域发电量为0kw，可以发电的区域根据光照，地形等给出了每平方公里年发电量x千瓦。
我们希望能够找到其中集中的矩形区域建设电站，能够获得良好的收益。

输入描述
第一行输入为调研的地区长，宽，以及准备建设的电站【长宽相等，为正方形】的边长最低要求的发电量
之后每行为调研区域每平方公里的发电量

输出描述
输出为这样的区域有多少个
*/

/*
*输入：
2 5 2 6
1 3 4 5 8
2 3 6 7 1

输出：
4
*/
func main6() {
	//凑成二位数组
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrP := strings.Fields(scanner.Text())

	m, _ := strconv.Atoi(arrP[0])
	n, _ := strconv.Atoi(arrP[1])
	//s, _ := strconv.Atoi(arrP[2])
	//m, _ := strconv.Atoi(arrP[3])

	matrix := make([][]int, m)

	for i := 0; i < m; i++ {
		scanner.Scan()
		arr := strings.Fields(scanner.Text())
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			v, _ := strconv.Atoi(arr[j])
			matrix[i][j] = v
		}
	}

}
