package practice

type NumArray []int

func ConstructorSumRange(nums []int) NumArray {
	s := make(NumArray, len(nums)+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	return s
}

// SumRange 思路:前缀和

func (s NumArray) SumRange(left, right int) int {
	return s[right+1] - s[left]
}
