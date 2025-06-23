package calculate

/*
*
用位运算实现加法
*/
func add(a int, b int) int {
	for b != 0 {
		// 无进位的和
		sum := a ^ b
		// 进位
		carry := (a & b) << 1
		a = sum
		b = carry
	}
	return a
}

func minus(a int, b int) int {
	return add(a, add(^b, 1))
}

/*
*
0110011
0001010
*/
func multiply(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 && b < 0 {
		return multiply(add(^a, 1), add(^b, 1))
	}
	if a < 0 {
		return add(^multiply(add(^a, 1), b), 1)
	}
	if b < 0 {
		return add(^multiply(a, add(^b, 1)), 1)
	}
	res := 0
	for i := 0; i < b; i++ {
		res = add(res, a)
	}
	return res
}
