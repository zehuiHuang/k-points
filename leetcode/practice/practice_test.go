package practice

import (
	"fmt"
	"testing"
)

func TestDecodeString(t *testing.T) {
	//abccdcdcdxyz
	//fmt.Println(decodeString2("abc3[cd]xyz"))
	//b3[a[2[c]]]
	fmt.Println(decodeString2("3[a2[c]]"))
}
