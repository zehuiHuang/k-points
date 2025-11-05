package practice

/**
思路:1先用切片模拟栈,包括进的栈和出的栈
2、每次pop或者peek时,都是从出的栈取值的(出的栈的数据都是从进栈来的)
3、出的栈按照先进先出规则,进的栈按照先进后出的规则
4、由于golang的栈是切片模拟的,所有看着可能奇怪
*/

type MyQueue struct {
	//用切片模拟栈
	//栈特点:先进后出
	//队列特点:先进先出
	inStack, outStack []int
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.outStack) == 0 && len(q.inStack) == 0
}
