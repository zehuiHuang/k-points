package practice

/*
*
思路:简单算法和KMP算法
*/
//简答算法
func strStr(haystack string, needle string) int {
	//ABCABCABD 9
	//ABCABD 6
	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		//从0到m-n,m-n后面不符合了
		a := i
		b := 0
		for b < n && haystack[a] == needle[b] {
			a++
			b++
		}
		if b == n {
			return i
		}
	}
	return -1
}

/*
*
KMP算法
思路:
1. 构建next数组
*/
func strStr2(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	nextArr := next(needle)
	//sCur表示子字串的指针,pCur表示模式串的指针
	sCur, pCur := 0, 0
	for sCur < m && pCur < n {
		if haystack[sCur] == needle[pCur] {
			sCur++
			pCur++
		} else if pCur > 0 {
			//匹配失败,根据next回推
			pCur = nextArr[pCur]
		} else {
			sCur++
		}
	}
	if pCur == n {
		return sCur - pCur
	} else {
		return -1
	}
}

//前后缀数组
/**
案例1
A  B C A B A T
-1 0 0 0 1 2 1
     j       i
C!=A

案例2
0 1 2 3 4 5 6
a a b a a f t
-1 0 1 0 1 2 0
     j     i

j=2,i=5
b!=f
即: needle[j]!=needle[i],则j=next[j],也就是j=next[2]=1
判断needle[1]和判断needle[5]是否相等


next[i]表示从0到i的字符的前后缀 相同的字符最大个数(对应字符串不包括needle[i])
*/
func next(needle string) []int {
	n := len(needle)
	nextArr := make([]int, n)
	i := 0
	j := -1
	nextArr[0] = -1
	//最后一位不处理
	for i < n-1 {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			nextArr[i] = j
		} else {
			//关键理解

			/**

			0 1 2 3 4 5 6
			a a b a a f t
			-1 0 1 0 1 2 0
			     j     i
			j=2,i=5
			当走到j=2,i=5时,由于b!=f 那么从那开始匹配呢?j肯定不能直接从0开始,nextArr[j]是什么呢?
			*/
			j = nextArr[j]
		}
	}
	return nextArr
}
