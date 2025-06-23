package base

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	hp := &RectHeap{}
	for i := 2; i < 3; i++ {
		*hp = append(*hp, Rectangle{i, i})
	}

	fmt.Println("原始slice: ", hp)

	// 堆操作
	heap.Init(hp)
	heap.Push(hp, Rectangle{100, 10})
	heap.Push(hp, Rectangle{1, 1})
	heap.Push(hp, Rectangle{1, 2})
	heap.Push(hp, Rectangle{10, 10})
	heap.Push(hp, Rectangle{100, 9})
	fmt.Println("top元素：", (*hp)[0])
	heap.Pop(hp)
	fmt.Println("top元素：", (*hp)[0])
	heap.Pop(hp)
	fmt.Println("top元素：", (*hp)[0])
	heap.Pop(hp)
	fmt.Println("top元素：", (*hp)[0])
	heap.Pop(hp)
	fmt.Println("top元素：", (*hp)[0])
	fmt.Println("最终slice: ", hp)
}
