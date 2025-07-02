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

/*
84. 柱状图中最大的矩形
思路：找到数组中左右两边分别第一个小于当前值的元素,比如：[2，1,5,6,2,3]中找到5的左右两边第一个小于他的值分别为1和2
*/
func largestRectangleArea(heights []int) int {
	ret := 0
	//单调栈
	stack := []int{}
	//将数组两边各加上0，保证每个数组元素一定能找到左右两边小于当前值的值
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			//弹出栈，找到栈中栈顶左边第一个小于栈顶的元素下标
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			ret = max(ret, (i-left-1)*heights[top])
		}
		stack = append(stack, i)
	}
	return ret
}
