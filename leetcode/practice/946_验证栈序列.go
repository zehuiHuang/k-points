package practice

// 3124
//
// pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
func validateStackSequences(pushed []int, popped []int) bool {
	//思路:模拟一个压栈和出栈的过程,如果最后栈是空的,说明符合栈先入后出规则
	stack := []int{}
	j := 0
	for _, v := range pushed {
		//先入栈
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
