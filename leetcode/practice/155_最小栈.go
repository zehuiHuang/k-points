package practice

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/**
思路:额外维护一个最小栈,保证栈顶是整个栈元素的最小值:
有一点需要注意,最小栈和常规栈的唯一区别是:常规栈实际存储值,而最小栈每次存储的一直都存的是最小值(最小栈和常规栈存储的数量相同)
*/

type MinStack struct {
	//最小栈.保证最小的在栈顶
	Stack []int
	//实际存储的值
	MStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	//添加时,判断最小栈栈顶是否小于该值
	if len(this.MStack) == 0 {
		this.MStack = append(this.MStack, val)
	} else {
		//获取栈顶最小元素
		m := this.MStack[len(this.MStack)-1]
		if m > val {
			this.MStack = append(this.MStack, val)
		} else {
			this.MStack = append(this.MStack, m)
		}
	}
}

func (this *MinStack) Pop() {
	//弹出栈顶元素
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MStack = this.MStack[:len(this.MStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MStack[len(this.MStack)-1]
}
