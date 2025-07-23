package main

//分披萨
/*
*
"吃货"和"馋嘴"两人到披萨店点了一份铁盘（圆形）披萨，并嘱咐店员将披萨按放射状切成大小相同的偶数个小块。但是粗心的服务员将披萨切成了每块大小都完全不同奇数块，且肉眼能分辨出大小。
由于两人都想吃到最多的披萨，他们商量了一个他们认为公平的分法：从"吃货"开始，轮流取披萨。除了第一块披萨可以任意选取外，其他都必须从缺口开始选。
他俩选披萨的思路不同。"馋嘴"每次都会选最大块的披萨，而且"吃货"知道"馋嘴"的想法。
已知披萨小块的数量以及每块的大小，求"吃货"能分得的最大的披萨大小的总和。
*/

/*
*
输入：
5
8
2
10
5
7
输出：19
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n int      // 披萨的数量
var arr []int  // 每块披萨的美味值
var dp [][]int // 记忆化数组，用于存储已计算过的状态

func main5() {
	//思路：定义馋嘴从披萨的两端（L,R）选择批量时,吃货能吃到的最到批量的量为：dp[L][R]
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	total, _ := strconv.Atoi(scanner.Text())

	arr = make([]int, total)
	for i := 0; i < total; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		arr[i] = v
	}
	n := len(arr)

	// 初始化记忆化数组
	dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 // 初始化为-1表示未计算
		}
	}

	ans := 0 // 初始化最大美味值
	// 遍历每块披萨，尝试以每块披萨作为起点
	for i := 0; i < n; i++ {
		L := (i + 1) % n     // 左边界：当前披萨的下一块
		R := (i + n - 1) % n // 右边界：当前披萨的前一块
		// 更新最大美味值：当前披萨值 + 剩余披萨的最优解
		current := allocation(L, R) + arr[i]
		if current > ans {
			ans = current
		}
	}

	fmt.Println(ans) // 输出最多能吃到的披萨的美味值总和
}

// 表示馋嘴从L到R按照贪心选择时，吃货能吃到的最多披萨的总理最多为dp[L][R]
func allocation(L, R int) int {
	// 如果当前状态已经计算过，则直接返回结果
	if dp[L][R] != -1 {
		return dp[L][R]
	}

	// 模拟对手的贪心选择：总是选择当前两端中较大的披萨
	if arr[L] > arr[R] {
		L = (L + 1) % n // 对手选择左端披萨，左边界右移
	} else {
		R = (R + n - 1) % n // 对手选择右端披萨，右边界左移
	}

	// 处理剩余披萨
	if L == R {
		dp[L][R] = arr[L] // 只剩一块披萨，直接返回值(因为是奇数个，且是吃货先选的，所有最后一个肯定是吃货选)
	} else {
		// 计算选择左端或右端披萨的最优解
		option1 := arr[L] + allocation((L+1)%n, R)   // 选择左端披萨
		option2 := arr[R] + allocation(L, (R+n-1)%n) // 选择右端披萨
		// 取两种选择的最大值
		if option1 > option2 {
			dp[L][R] = option1
		} else {
			dp[L][R] = option2
		}
	}

	return dp[L][R]
}
