package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 流浪地球
/**
流浪地球计划在赤道上均匀部署了N个转向发动机，按位置顺序编号为0~N-1。

初始状态下所有的发动机都是未启动状态;
发动机启动的方式分为”手动启动"和”关联启动"两种方式;
如果在时刻1一个发动机被启动，下一个时刻2与之相邻的两个发动机就会被”关联启动”;
如果准备启动某个发动机时，它已经被启动了，则什么都不用做;
发动机0与发动机N-1是相邻的;
地球联合政府准备挑选某些发动机在某些时刻进行“手动启动”。当然最终所有的发动机都会被启动。

哪些发动机最晚被启动呢?
*/
/**
输入描述
第一行两个数字N和E，中间有空格
N代表部署发动机的总个数，E代表计划手动启动的发动机总个数
1<N<=1000,1<=E<=1000,E<=N
接下来共E行，每行都是两个数字T和P，中间有空格
T代表发动机的手动启动时刻，P代表此发动机的位置编号。
0<=T<=N.0<=P<N
输出描述
第一行一个数字N，以回车结束
N代表最后被启动的发动机个数
第二行N个数字，中间有空格，以回车结束
每个数字代表发动机的位置编号，从小到大排序
*/

/*
*
输入：
8 2
0 2
0 6

输出：
2
0 4
*/
func main26() {
	//思路：将数组存储为执行时间
	//1、将所有发动机组装成一个数组，里面的值初始化为-1
	//2、将手动开启的引擎对应下标的值设为启动时间，并找到启动时间的最小值
	//3、循环遍历数组，找到启动时间最小对应的索引位置，对其左边和右边的值设置为time+1
	//4、遍历过程中若不符合，则time+1重新对数组进行遍历
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(input1[0])
	k, _ := strconv.Atoi(input1[1])
	arr := make([]int, n)

	//初始化为-1
	for i := range arr {
		arr[i] = -1
	}

	//将手动开启的下标设置为开始的时间time
	minTime := math.MaxInt64
	for i := 0; i < k; i++ {
		scanner.Scan()
		points := strings.Fields(scanner.Text())
		time, _ := strconv.Atoi(points[0])
		index, _ := strconv.Atoi(points[1])
		arr[index] = time
		minTime = min(minTime, time)
	}

	updateFlag := true
	for updateFlag {
		for i := range arr {
			if arr[i] == minTime {
				//将左右两边的都设置为minTime+1
				left := (i - 1 + n) % n
				right := (i + 1) % n
				if arr[left] < 0 {
					arr[left] = minTime + 1
				}
				if arr[right] < 0 {
					arr[right] = minTime + 1
				}
			}
		}
		minTime++
		//判定是否已经都启动完
		updateFlag = false
		for i := range arr {
			if arr[i] == -1 {
				updateFlag = true
			}
		}
	}
	//过滤出来找到最大值
	count := 0
	idx := []string{}

	maxV := -1
	for i := range arr {
		maxV = max2(maxV, arr[i])
	}

	for i := range arr {
		if arr[i] == maxV {
			count++
			idx = append(idx, strconv.Itoa(i))
		}
	}
	fmt.Println(count)
	fmt.Println(strings.Join(idx, " "))
}
