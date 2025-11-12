package practice

//[1,3,-1,-3,5,3,6,7],k=3
//[5,3,-1,-3,5,3,6,7],k=3

// 思路:关键点:滑动窗口范围内,如何挑选最大值问题

//使用单调栈结构,单调栈(最小栈,保证栈顶是最小值,然后从栈底取最大值)

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	//为了判定 某值经过滑动后是否还在窗口内,栈内存的是索引地址index
	stack := []int{}
	//push时栈顶保持最小值
	var push func(i int)
	push = func(i int) {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, nums[stack[0]])
	for i := k; i < len(nums); i++ {
		//入栈
		push(i)
		//判定当前的最大值是否还在窗口内,如果不在责需要移除
		m := stack[0]
		if m < i-k+1 {
			stack = stack[1:]
		}
		ans = append(ans, nums[stack[0]])
	}
	return ans
}
