package practice

/**
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：

输入：x = 2.10000, n = 3
输出：9.26100
示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
*/

/**
思路:使用快速密,也就是x的n次方中的n 转化为二进制算法
比如 n=8 ,二进制就是1000 ,那么还等于0*1+0*2+0*4+1*8=8
那么对n进行二分,内部变成乘2
例如:x的4次方可以等价于x的2次方的值的二次方 ,那么n就减少了一半
同时二分后分为两种情况,一种是能被整除,一种是余数为1
针对余数为1的可以将多出来的直接向乘即可

*/

func myPow(x float64, n int) float64 {
	if x == 1.0 {
		return x
	}
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n > 0 {
		//为奇数时,多的那个1直接乘进去
		if n&1 == 1 {
			res = res * x
		}
		//x变成了x的平方
		x = x * x
		//将n右移一位,即除以2
		n = n >> 1
	}
	return res
}

// 时间超时,复杂度过高
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return x
	}
	res := x
	if n > 0 {
		for i := 1; i < n; i++ {
			res = res * x
		}
	} else if n < 0 {
		for i := 1; i < -n; i++ {
			res = res * x
		}
		res = 1 / res
	}
	return res
}
