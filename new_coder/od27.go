package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 任务总执行时长

/*
*
任务编排服务负责对任务进行组合调度。参与编排的任务有两种类型，其中一种执行时长为taskA，另一种执行时长为taskB。
任务一旦开始执行不能被打断，且任务可连续执行。服务每次可以编排num个任务。请编写一个方法，生成每次编排后的任务所有可能的总执行时长
*/

/*
*
输入：
1,2,3
输出：
[3, 4, 5, 6]
*/
func main27() {
	//思路：排列组合问题，将taskA*i+taskB*(count-i) 每一种情况都计算一遍，count为从0～num
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	taskA, _ := strconv.Atoi(input[0])
	taskB, _ := strconv.Atoi(input[1])
	count, _ := strconv.Atoi(input[2])
	if count == 0 {
		fmt.Println("[]")
		return
	}
	mp := make(map[string]string)
	for i := 0; i <= count; i++ {
		key := strconv.Itoa(taskA*i + taskB*(count-i))
		mp[key] = key
	}
	ans := make([]int, len(mp))
	i := 0
	for key := range mp {
		ans[i], _ = strconv.Atoi(key)
		i++
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	fmt.Println(ans)

}
