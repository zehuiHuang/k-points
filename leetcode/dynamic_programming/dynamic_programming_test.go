package dynamic_programming

import (
	"fmt"
	"math"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	result := lastStoneWeight(stones)
	fmt.Println(result)
}

func TestMaxSubArray(t *testing.T) {
	//nums := []int{-10, -100}
	//println(maxSubArray(nums))
	fmt.Println(math.MaxInt64)
}

func TestName(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 3
	coinChange2(coins, amount)
}
