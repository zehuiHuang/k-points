package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 幼儿园分班问题

/*
*
题目描述
儿园两个班的小朋友在排队时混在了一起，每位小朋友都知道自己是否与前面一位小朋友同班，请你帮忙把同班的小朋友找出来。

小朋友的编号是整数，与前一位小朋友同班用Y表示，不同班用N表示。

输入描述
输入为空格分开的小朋友编号和是否同班标志。

比如：6/N 2/Y 3/N 4/Y，表示4位小朋友，2和6同班，3和2不同班，4和3同班。

其中，小朋友总数不超过999，每个小朋友编号大于0，小于等于999。

不考虑输入格式错误问题。

输出描述
输出为两行，每一行记录一个班小朋友的编号，编号用空格分开，且：

编号需按照大小升序排列，分班记录中第一个编号小的排在第一行。
若只有一个班的小朋友，第二行为空行。
若输入不符合要求，则直接输出字符串ERROR。
*/

/*
*
输入：
1/N 2/Y 3/N 4/Y
或
1/N 2/Y 3/N 4/Y 5/Y
输出：
1 2
3 4
*/

func main31() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	nums := strings.Fields(line)

	// 处理第一个学生
	first := strings.Split(nums[0], "/")
	classA := []string{first[0]}
	classB := []string{}
	temp := [][]string{classA, classB} // temp[0]是classA, temp[1]是classB

	// 处理后续学生
	for i := 1; i < len(nums); i++ {
		parts := strings.Split(nums[i], "/")
		if len(parts) < 2 {
			continue
		}

		id_ := parts[0]
		f := parts[1]

		if f == "N" {
			// 反转temp中的两个班级顺序
			temp[0], temp[1] = temp[1], temp[0]
		}

		// 将学生添加到当前第一个班级
		temp[0] = append(temp[0], id_)
	}

	for i := 0; i < len(temp); i++ {
		sort.Slice(temp[i], func(x, y int) bool {
			return temp[i][x] < temp[i][y]
		})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][0] < temp[j][0]
	})

	for i := range temp {
		fmt.Println(strings.Join(temp[i], " "))
	}
}
