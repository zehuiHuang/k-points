package stack

//单调栈

/*
739. 每日温度：https://leetcode.cn/problems/daily-temperatures/description/
思路：使用单调栈，栈从上到下为单调递增(目的是为了获取当前i比最小的大多少，然后更新栈顶的差值，最后将栈顶弹出)
*/

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	//定义单调栈
	stack := []int{}

	for i := 0; i < length; i++ {
		//获取当前遍历的值,并和栈顶的值进行对比（栈里存储的是数组下标）
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			//获取栈顶的数据（温度的数组下标）
			index := stack[len(stack)-1]
			//弹出栈顶数据
			stack = stack[:len(stack)-1]
			ans[index] = i - index
		}
		stack = append(stack, i)
	}
	return ans
}
