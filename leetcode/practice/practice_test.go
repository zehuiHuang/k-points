package practice

import (
	"fmt"
	"testing"
	"time"
)

func TestDecodeString(t *testing.T) {
	//abccdcdcdxyz
	//fmt.Println(decodeString2("abc3[cd]xyz"))
	//b3[a[2[c]]]
	fmt.Println(decodeString2("3[a2[c]]"))
}

func TestSetZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
}

func TestName(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(1 << 4)
}

func TestNext(t *testing.T) {
	fmt.Println(next("aabaaft"))
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true,
		},
		{
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			expected: true,
		},
	}

	for i, test := range tests {
		result := wordBreak(test.s, test.wordDict)
		if result != test.expected {
			t.Errorf("Test case %d failed: wordBreak(%q, %v) = %v, expected %v",
				i, test.s, test.wordDict, result, test.expected)
		}

		result2 := wordBreakOptimized(test.s, test.wordDict)
		if result2 != test.expected {
			t.Errorf("Test case %d failed: wordBreakOptimized(%q, %v) = %v, expected %v",
				i, test.s, test.wordDict, result2, test.expected)
		}
	}
}

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
		desc     string
	}{
		{
			nums:     []int{1, 5, 11, 5},
			expected: true, // 可分为[1,5,5]和[11]
			desc:     "可等分数组",
		},
		{
			nums:     []int{1, 2, 3, 5},
			expected: false, // 无法等分
			desc:     "不可等分数组",
		},
	}

	for i, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)

		result := canPartition(numsCopy)
		if result != test.expected {
			t.Errorf("Test case %d (%s) failed: canPartition(%v) = %t, expected %t",
				i, test.desc, test.nums, result, test.expected)
		}
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
		desc     string
	}{
		{
			coins:    []int{1, 3, 4},
			amount:   6,
			expected: 2, // 3+3
			desc:     "示例硬币",
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1, // 无法凑成3
			desc:     "无法凑成金额",
		},
	}

	for i, test := range tests {
		coinsCopy := make([]int, len(test.coins))
		copy(coinsCopy, test.coins)

		result := coinChange(coinsCopy, test.amount)
		if result != test.expected {
			t.Errorf("Test case %d (%s) failed: coinChange(%v, %d) = %d, expected %d",
				i, test.desc, test.coins, test.amount, result, test.expected)
		}
	}
}

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
		desc     string
	}{
		{
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5, // 有5种方法添加+/-符号使得和等于3
			desc:     "目标和问题示例",
		},
		{
			nums:     []int{1},
			target:   1,
			expected: 1, // 只能是+1
			desc:     "单元素正目标",
		},
	}

	for i, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)

		result := findTargetSumWays(numsCopy, test.target)
		if result != test.expected {
			t.Errorf("Test case %d (%s) failed: findTargetSumWays(%v, %d) = %d, expected %d",
				i, test.desc, test.nums, test.target, result, test.expected)
		}
	}
}
