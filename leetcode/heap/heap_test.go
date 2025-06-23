package heap

import (
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	maxSlidingWindow([]int{5, 3, -1, -3, 5, 3, 6, 7}, 3)
}
