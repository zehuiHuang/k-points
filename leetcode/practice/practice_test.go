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
