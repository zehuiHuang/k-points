package practice

// s="ace"
// t="abcde"
// 思路:双指针,字符相等,则同时移动s和t的指针位置,直到s的指针移动到末尾
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	sIndex := 0
	for tIndex := 0; tIndex < len(t); tIndex++ {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		if sIndex == len(s) {
			return true
		}
	}
	return false
}
