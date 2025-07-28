package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 贪心歌手

/*
*
题目描述
一个歌手准备从A城去B城参加演出。

按照合同，他必须在 T 天内赶到
歌手途经 N 座城市
歌手不能往回走
每两座城市之间需要的天数都可以提前获知。
歌手在每座城市都可以在路边卖唱赚钱。

经过调研，歌手提前获知了每座城市卖唱的收入预期：
如果在一座城市第一天卖唱可以赚M，后续每天的收入会减少D（第二天赚的钱是 M - D，第三天是 M - 2D ...）。如果收入减少到 0 就不会再少了。
歌手到达后的第二天才能开始卖唱。如果今天卖过唱，第二天才能出发。
贪心的歌手最多可以赚多少钱？

输入描述
第一行两个数字 T 和 N，中间用空格隔开。

T 代表总天数，0 < T < 1000
N 代表路上经过 N 座城市，0 < N < 100
第二行 N+1 个数字，中间用空格隔开。代表每两座城市之间耗费的时间。

其总和 ≤ T。
接下来 N 行，每行两个数字 M 和 D，中间用空格隔开。代表每个城市的输入预期。

0 < M < 1000
0 < D < 100
输出描述
一个数字。代表歌手最多可以赚多少钱。以回车结束。
*/

/*
*
输入:
10 2
1 1 2
120 20
90 10
输出:
540
*/

// 定义最小堆类型
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main43() {
	//思路：将每天能赚到的钱都放入优先级队列中(小根堆)
	//1、前期只要当前还能赚钱且没有用完卖唱的总时间，就要放入队列中
	//2、如果后续发现卖唱的总时间不够了，那么就需要比较当前当天卖唱的钱和队列中的top比较，如果发现没top多，那么当前城市的这边没必要待下去，
	// 如果发现当天的卖唱钱比top的多，那么堆里之前放入的那天就可以剔除掉，把当天卖唱的钱放入队列中
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input1 := strings.Fields(scanner.Text())
	t, _ := strconv.Atoi(input1[0])
	n, _ := strconv.Atoi(input1[1])

	scanner.Scan()
	times := strings.Fields(scanner.Text())
	//由于间隔n个城市，所有城市之间的距离数量为n+1个
	costTime := 0

	for i := range times {
		t, _ := strconv.Atoi(times[i])
		costTime += t
	}

	usedTotal := t - costTime

	mds := make([][2]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		abc := strings.Fields(scanner.Text())
		m, _ := strconv.Atoi(abc[0])
		d, _ := strconv.Atoi(abc[1])
		mds = append(mds, [2]int{m, d})
	}

	h := &MinHeap{}
	heap.Init(h)

	for i := range mds {
		money := mds[i][0]
		reduce := mds[i][1]
		for money > 0 {
			//如果小堆中已经存储了剩余的天数，那么就对比当前挣到的钱和历史最少挣的钱做对比
			if h.Len() >= usedTotal {
				//获取堆最小元素
				top := (*h)[0]
				if top > money { //历史挣的最少钱还比当前的多，则这天该城市也没必要待下去了
					break
				} else {
					//如果当天能挣到的比之前最少的要多，那么堆的top弹出，将当天挣的钱push进去
					heap.Pop(h)
				}
			}
			heap.Push(h, money)
			money -= reduce
		}
	}
	totalMoney := 0
	for i := range *h {
		totalMoney += (*h)[i]
	}
	fmt.Println(totalMoney)
}
