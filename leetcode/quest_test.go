package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	num := []int{100, 4, 200, 1, 3, 2}
	print(LongestConsecutive2(num))
}

func TestJump(t *testing.T) {
	num := []int{2, 3, 1, 1, 4}
	print(jump(num))
	//num := []int{3, 2, 1, 0, 4}
	//num := []int{2, 3, 1, 1, 4}
	//num := []int{2, 3, 1, 4, 0, 5}
	//print(canJump(num))
}

func TestCanCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	print(canCompleteCircuit(gas, cost))
}

func TestFindKthLargestt(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	print(findKthLargest(nums, k))
}

func TestAbc(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}
	copy(s2, s1)
	fmt.Println(s1) //输出:[1 2]
	fmt.Println(s2) //输出:[1 2 5 6]
}

func TestMerge(t *testing.T) {
	fmt.Println(merge([]int{1, 2, 4, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupAnagrams(strs)
}

func TestSubarraySum2(t *testing.T) {
	nums := []int{1, 2, 3}
	subarraySum2(nums, 3)
}
