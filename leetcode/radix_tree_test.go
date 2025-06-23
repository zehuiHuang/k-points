package leetcode

import "testing"

func TestRadixInsert(t *testing.T) {
	radix := NewRadix()
	//情况1
	//radix.Insert("abc")
	//radix.Insert("abd")
	//radix.Insert("af")

	radix.Insert("abc")
	radix.Insert("abdf")
	radix.Insert("abde")

}
