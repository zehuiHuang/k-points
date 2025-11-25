package heap

import (
	"fmt"
	"sort"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	maxSlidingWindow([]int{5, 3, -1, -3, 5, 3, 6, 7}, 3)
}

func TestName(t *testing.T) {
	a := []int{2, 3, 7, 1, 6, 2}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}
