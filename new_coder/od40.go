package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 士兵过河

/*
*
题目描述
一支N个士兵的军队正在趁夜色逃亡，途中遇到一条湍急的大河。
敌军在T的时长后到达河面，没到过对岸的士兵都会被消灭。
现在军队只找到了1只小船，这船最多能同时坐上2个士兵。

当1个士兵划船过河，用时为 a[i]；0 <= i < N
当2个士兵坐船同时划船过河时，用时为max(a[j],a[i])两士兵中用时最长的。
当2个士兵坐船1个士兵划船时，用时为 a[i]*10；a[i]为划船士兵用时。
如果士兵下河游泳，则会被湍急水流直接带走，算作死亡。
请帮忙给出一种解决方案，保证存活的士兵最多，且过河用时最短。

输入描述
第一行：N 表示士兵数(0<N<1,000,000)
第二行：T 表示敌军到达时长(0 < T < 100,000,000)
第三行：a[0] a[1] … a[i]… a[N- 1]
a[i]表示每个士兵的过河时长。
(10 < a[i]< 100; 0<= i< N）

输出描述
第一行：”最多存活士兵数” “最短用时”

备注
1）两个士兵的同时划船时，如果划速不同则会导致船原地转圈圈；所以为保持两个士兵划速相同，则需要向划的慢的士兵看齐。
2）两个士兵坐船时，重量增加吃水加深，水的阻力增大；同样的力量划船速度会变慢；
3）由于河水湍急大量的力用来抵消水流的阻力，所以2）中过河用时不是a[i] *2，
而是a[i] * 10。
*/

/*
*输入
5
43
12 13 15 20 50
输出	3 40
说明	可以达到或小于43的一种方案：
第一步：a[0] a[1] 过河用时：13
第二步：a[0] 返回用时：12
第三步：a[0] a[2] 过河用时：15
*/
func main40() {
	//二分法，假设将N/2个士兵运到河对岸花费time时间，则和时间限制limit做对比，
	//若time<limit,则说明在time时间内可以将n/2个所有士兵运过去，那么可以尝试n/2+1个士兵花费的最小时间，然后和limit多对比，以此循环
	//若time>limit,则说明不能n/2个士兵会运输不完，则应该n/2-1在尝试和循环
	//若time=limit,则说明刚好在限制时间内运走N/2个士兵
	//最后关键是是计算指定某些人数最多可用时间

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	limit, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	arr := strings.Fields(scanner.Text())

	tables := make([]int, len(arr))

	for i := range arr {
		tables[i], _ = strconv.Atoi(arr[i])
	}
	//升序排列
	sort.Slice(tables, func(i, j int) bool {
		return tables[i] < tables[j]
	})

	var getTime func(n int, times []int) int

	getTime = func(n int, times []int) int {
		cost := 0
		for n > 0 {
			if n == 1 {
				cost += times[0]
				break
			} else if n == 2 {
				cost += times[1]
				break
			} else if n == 3 {
				cost += times[2] + times[0] + times[1]
				break
			} else {
				//两种做法：
				//1、每次都用划的最快的times[0]作为每次的来回运输人
				//2、times[0]和times[1]先过去，回来的时times[0],然后将最慢的和times[1]运送过去，回来times[1]
				cost += min(times[len(times)-1]+times[0]+times[len(times)-2]+times[0],
					times[1]+times[0]+times[len(times)-1]+times[1])
			}
			n -= 2
		}

		return cost
	}
	//通过二分查找
	min := 0
	max := N
	ans := ""
	for min <= max {
		mid := (min + max) / 2
		cpArr := tables[:mid]
		time := getTime(mid, cpArr)
		//如果占用的时间比限制时间还打，说明限制时间内不能将所有时间全部运送过去
		if time > limit {
			max = mid - 1
		} else if time < limit {
			ans = strconv.Itoa(mid) + " " + strconv.Itoa(time)
			min = mid + 1
		} else {
			//相等，则说明正好可以在时间limit下将所有士兵运完
			ans = strconv.Itoa(mid) + " " + strconv.Itoa(time)
			break
		}
	}
	fmt.Println(ans)
}
