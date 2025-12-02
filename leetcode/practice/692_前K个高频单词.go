package practice

import "container/heap"

type pair struct {
	w string
	c int
}
type hp2 []pair

func (h hp2) Len() int {
	return len(h)
}
func (h hp2) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.c < b.c || a.c == b.c && a.w > b.w
}
func (h hp2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp2) Push(v interface{}) {
	*h = append(*h, v.(pair))
}
func (h *hp2) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp2{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
