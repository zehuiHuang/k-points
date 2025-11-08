package leetcode

// 242. 有效的字母异位词
func isAnagram(s string, t string) bool {
	l1 := len(s)
	l2 := len(t)
	m1 := make(map[rune]int)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		m1[rune(s[i])] += 1
	}
	for i := 0; i < l2; i++ {
		m1[rune(t[i])] -= 1
	}
	for _, v := range m1 {
		if v != 0 {
			return false
		}
	}
	return true
}

// 387. 字符串中的第一个唯一字符
/**
思路:使用map,存储每个字符出现的次数,然后遍历字符串,第一个为1的即为答案
*/
func firstUniqChar(s string) int {
	mp := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		mp[rune(s[i])] += 1
	}
	for i := 0; i < len(s); i++ {
		if mp[rune(s[i])] == 1 {
			return i
		}
	}
	return -1
}

// 205. 同构字符串
// 思路:对两个字符串中的字符分别进行映射
// 比如aab和cca:a->c,a->c,b->a  这种key、value映射关系一旦确定,同构字符串的字符关系映射就不会变,如果变了,则表示不是同构字符串
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp1 := make(map[rune]rune)
	mp2 := make(map[rune]rune)
	n := len(s)
	for i := 0; i < n; i++ {
		mp1[rune(s[i])] = rune(t[i])
		mp2[rune(t[i])] = rune(s[i])
	}
	for i := 0; i < n; i++ {
		if mp1[rune(s[i])] != 0 && mp1[rune(s[i])] != rune(t[i]) ||
			mp2[rune(t[i])] != 0 && mp2[rune(t[i])] != rune(s[i]) {
			return false
		}
	}
	return true
}

// 205. 回文排序
// 思路:统计字符奇偶的个数,如果奇数的大于1则不能组合成回文
func canPermutePalindrome(s string) bool {
	mp := make(map[rune]int)
	for _, v := range s {
		mp[rune(v)] += 1
	}
	odd := 0
	for _, v := range mp {
		if v%2 == 1 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}
	return true
}

// 409. 最长回文串
// 思路,异构的回文传,那么只要判定奇数的字符串个数为0个或个就好
func longestPalindrome(s string) int {
	n := len(s)
	mp := make(map[rune]int)
	odd := 0
	res := 0
	for i := 0; i < n; i++ {
		mp[rune(s[i])] += 1
	}
	for _, v := range mp {
		rem := v % 2 //判断是否是偶数
		//该步骤的目的是让奇数个字母变成偶数个
		res += v - rem
		//奇数个只留一个放到中间即可
		if rem == 1 {
			odd = 1
		}
	}
	return odd + res
}
