package practice

import "strconv"

/*
*
给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位上的数字。
*/
func findNthDigit(n int) int {
	//2 * 10 * 9=180
	//n=127 126 125 124 123 122 121 120 119 118
	//n=127-9=118

	/**
	找规律:
	1-9        首位是1      位数位1    数字数量为9    数位数量为1*9
	10-99      start=10    digit=2   count=90        2*90
	100-999    start=100   digit=3   count=900       3*900
	...n       start=1..   digit=n   count=9*start   9*start*digit
	*/

	//三步走:1确定n所在的位数,2确定n所在的数字,3确定n是num中的哪一位
	//例如n=205 求第1234567891011121314....第205个位置的数字是几
	digit, start, count := 1, 1, 9
	for n > count {
		//第一轮 n=196,count= 2*90
		//第二轮 n=16,count=  3*900
		//第二轮之后就跳出循环,得出digit=3,start=100
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}

	//start=100,(205-1)/100=2
	//比如从100开始,101,102,102....
	//此时n已经被减到从100开始数,第多少个,此时n=16,除以digit=3是因为一个数占了3位
	nums := start + (n-1)/digit
	str := strconv.Itoa(nums)
	//取余表示在在整个字符串的下标,即某3位数的某一个下标索引
	index := (n - 1) % len(str)
	return int(str[index] - '0')
}
