package practice

import (
	"container/heap"
	"sort"
)

// 思路:一种为有序队列加 双指针,一种为结合大堆和小堆完成
// 以大堆和小堆为例:

type MedianFinder struct {
	left  hp // 入堆的元素取相反数，变成最大堆
	right hp // 最小堆
}

//func Constructor() MedianFinder {
//
//}

func (this *MedianFinder) AddNum(num int) {
	//1首先设定左堆和右堆,左边为最大堆(存对应值得负数),右边为最小堆,即右边的都比左边的大
	if this.left.Len() == this.right.Len() {
		//先放入右堆,并从右堆取出最小的, 并放入到左堆
		heap.Push(&this.right, num)
		v := -heap.Pop(&this.right).(int) //注意是负数
		heap.Push(&this.left, v)
	} else {
		//先放入左堆,然后从左堆取之后放入到右堆
		heap.Push(&this.left, -num)
		v := -heap.Pop(&this.left).(int)
		heap.Push(&this.right, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	//取之
	//如果两边长度相等,则各自取一个除以2即可,如果左边比右边大一,则取左边即为答案
	if this.left.Len() > this.right.Len() {
		return -float64(this.left.IntSlice[0])
	}
	return float64((this.right.IntSlice[0])-(this.left.IntSlice[0])) / 2.0
}

//最小堆

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
