package heap

import (
	"container/heap"
	"fmt"
	"sort"
)

// MinHeap 定义最小堆类型
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	//升序:小顶堆,小的在最上面
	//1,2,3,4
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

func findTopKMinHeap(nums []int, k int) []int {
	if k <= 0 || len(nums) < k {
		return nil
	}
	h := &MinHeap{}
	heap.Init(h)
	// 遍历数组维护堆大小
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	// 提取堆中元素（无序，若需有序可逆序）
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

// ------------------------------------------------------------
var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}
func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

/*
*滑动窗口最大值 leetcode:239
*方法：单调队列
 */
/*
思路，初始值向优先级队列放入窗口中的值，然后弹出，即为最大值，然后从k～n开始遍历（窗口向右移动）
此时最大值可能并不在窗口中，所以，弹出的值如果不在窗口中则丢弃，继续弹出，直到出现在窗口中
*/
func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	//初始化优先级队列
	q := &hp{make([]int, k)}
	//将窗口中的值放入优先级队列
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	//初始化优先级队列
	heap.Init(q)
	//遍历窗口，从k～n
	n := len(nums)
	//初始化结果集
	ans := make([]int, 1, n-k+1)
	//将窗口中的最大值放入结果集
	ans[0] = nums[q.IntSlice[0]]
	//遍历窗口，从k～n
	for i := k; i < n; i++ {
		//将当前值放入优先级队列
		heap.Push(q, i)
		//如果优先级队列中的最大值(索引值)不在窗口中，则弹出
		fmt.Println(q.IntSlice[0])
		//第一次：索引值要大于0才能放入结果（小于等于0的都要弹出），第二次：索引值要大于1才能放入结果（小于等于1的要弹出），
		//以此类推，将窗口中的最值值放入结果集
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		//将当前窗口中的最大值放入结果集
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
	//5,  3, -1, -3,  5,  3,  6,  7
	//0   1   2   3   4   5   6   7
}

//leetcode:239
func maxSlidingWindow2(nums []int, k int) []int {
	//单调栈（从左到右单调递减）
	q := []int{}
	//向单调栈放入元素，如果发现当前元素的值大于单调栈的栈顶元素的值（最小值），则遍历移除掉栈顶元素（我在你右边还比你大，那么在它左边的都可以去掉）
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

// 滑动窗口最大值：单调栈
func maxSlidingWindow3(nums []int, k int) []int {
	//nums = [1,3,-1,-3,5,3,6,7], k = 3
	n := len(nums)
	stack := []int{}
	ans := make([]int, 1, n-k+1)
	push := func(i int) {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, stack[0])
	for i := k; i < n-1; i++ {
		push(i)
		if stack[0] < i-k+1 {
			stack = stack[1:]
		}
		ans = append(ans, stack[0])
	}
	return ans
}

// 面试题 17.14. 最小K个数
/**
思路：最大堆，向堆初始化数据，然后从k～n，保证堆中的梳理为k，offer一个，就弹出一个，最后堆中剩下的即为最小k个数了
*/
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}
