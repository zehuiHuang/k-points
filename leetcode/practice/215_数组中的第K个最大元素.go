package practice

import (
	"container/heap"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

*/
// 思路1:最直接的将数据放入到小顶堆中,并保持小堆的大小为k,那么获取堆顶元素即为答案
//思路2:
func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	//初始化
	hp := &MinHeap{}
	heap.Init(hp)
	for i := 0; i < len(nums); i++ {
		heap.Push(hp, nums[i])
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}
	return heap.Pop(hp).(int)
}

//最大堆或最小堆的实现,若继承了sort.IntSlice则无需重写Len和Swap
/**
首先要定义堆类型
type MinHeap []int
第二要实现五个方法
Len()int  :长度
Swap(i,j int) :交换位置
Less(i,j int)bool :排序
Push(v interface{}) : 写入
Pop()interface{} :弹出
*/

type MinHeap []int

//	type hp2 struct {
//		sort.IntSlice
//	}
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	//小堆
	return h[i] > h[j]
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
