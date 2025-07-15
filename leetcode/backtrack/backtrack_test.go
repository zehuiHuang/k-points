package backtrack

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("abba"))
}

func TestName(t *testing.T) {
	path := []string{"a", "b", "c"}
	str := strings.Join(path, ".")
	fmt.Println(str)
}
