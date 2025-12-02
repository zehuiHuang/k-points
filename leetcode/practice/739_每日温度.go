package practice

/*
*思路:使用单调栈，栈从上到下为单调递增
1、只要右边不出现比当前值大的情况,就放入栈,栈放的是数组下标
2、当 当前值出现了比栈顶数(栈顶存的下标,要转化成值)大的值时,则栈顶的下标索引(也就是ans索引)对应的值为i-index
*/
func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		/**
		例如:4,3,2,1,5,判定4的ans,即下标为0的值
		那么栈会放入0,1,2,3,当遍历到下标为4时,他的值大于栈顶3的元素,所以下标3的ans值为i-index,结果为1,最后从栈弹出下标为3的值
		继续判断当前值是否大于栈顶元素,还是大于,则继续上面的逻辑
		*/
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			ans[index] = i - index
			//弹出栈顶
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
