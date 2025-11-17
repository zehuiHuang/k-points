package practice

// 思路: 模拟,整体思路和54题基本类似
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}
	v := 0
	up := 0
	down := n - 1 //---------行数
	left := 0
	right := n - 1 //---------列数
	for {
		//先从左向右,列数(从第一列,到最后一列)
		for i := left; i <= right; i++ {
			v += 1
			matrix[up][i] = v //行不变
		}
		if up == down {
			break
		} else {
			up++
		}

		//从上到下
		for i := up; i <= down; i++ {
			v += 1
			matrix[i][right] = v //列不变
		}
		if left == right {
			break
		} else {
			right--
		}

		//从右向左
		for i := right; i >= left; i-- {
			v += 1
			matrix[down][i] = v //行数不变
		}
		if up == down {
			break
		} else {
			down--
		}

		//从下到上
		for i := down; i >= up; i-- {
			v += 1
			matrix[i][left] = v //列数不变
		}
		if left == right {
			break
		} else {
			left++
		}

	}
	return matrix
}
