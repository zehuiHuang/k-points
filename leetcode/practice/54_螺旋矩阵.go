package practice

// 整体思路: 模拟过程,一个大周期为:左->右=》上->下 => 右->左 => 下->上
// 在这个过程中,逐渐缩小上和下、左和右的边界,直到重合
func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	m := len(matrix) - 1    //行索引
	n := len(matrix[0]) - 1 //列索引
	//定义上下左右边界
	left := 0
	right := n //从左到右,走的其实是列数
	up := 0
	down := m //从上到下,走的其实是行数
	for {
		//从左向右移动,
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		//如果在大层的遍历中(左->右=》上->下 => 右->左 => 下->上 为一个大周期)up已经加到了和行数相等,则说明已经遍历完了(上下边界重合了)
		if up == down {
			//若不满足,则直接执行完毕,直接跳出循环
			break
		} else {
			//若满足条件,则重新定义边界(加1是因为从左到右时,当前的一层已经被用过,行数减少了1)
			up++
		}

		//从上到下
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right]) //从上到下,涉及到行数
		}
		//如果在大层的遍历中(左->右=》上->下 => 右->左 => 下->上 为一个大周期)边界重合,则说明遍历完毕
		if right == left {
			break
		} else {
			//同理,若满足条件,则重新定义边界(减1是因为上到下时,当前的一列已经被用过,列数应减少了1)
			right--
		}

		//从右到左
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i]) //行不变
		}
		if up == down {
			break
		} else {
			down--
		}

		//从下到上
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left]) //列不变,
		}
		if left == right {
			break
		} else {
			left++
		}
	}
	return ans
}
