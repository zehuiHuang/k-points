package dynamic_programming

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(stones)
	fmt.Println(result)
}
