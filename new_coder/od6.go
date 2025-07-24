package main

import (
	"bufio"
	"fmt"
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
	//思路：在整个区域的范围内（在m*n的范围内s*s的移动）分别进行向右向下移动，每移动一层，便计算s*s 在整个范围内所占的面积
	//凑成二位数组
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arrP := strings.Fields(scanner.Text())
	//分别为地区的长、宽、电站边长、最低发电量
	m, _ := strconv.Atoi(arrP[0])   //2
	n, _ := strconv.Atoi(arrP[1])   //5
	s, _ := strconv.Atoi(arrP[2])   //2
	min, _ := strconv.Atoi(arrP[3]) //6

	matrix := make([][]int, m)
	//2 5 2 6
	//1 3 4 5 8
	//2 3 6 7 1
	for i := 0; i < m; i++ {
		scanner.Scan()
		arr := strings.Fields(scanner.Text())
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			v, _ := strconv.Atoi(arr[j])
			matrix[i][j] = v
		}
	}

	ans := 0
	for i := s; i <= m; i++ {
		for j := s; j <= n; j++ {
			//在整个区域的范围内（在m*n的范围内s*s的移动）分别进行向右向下移动，
			square := 0
			for x := i - s; x < m; x++ {
				for y := j - s; y < n; y++ {
					square += matrix[x][y]
				}
			}
			if square >= min {
				ans++
			}
		}
	}
	fmt.Println(ans)

}
