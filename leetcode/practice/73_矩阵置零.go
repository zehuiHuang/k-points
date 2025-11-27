package practice

/**
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
*/
// 方法1:使用额外数组
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowHasZero := make([]bool, m) // 行是否包含 0
	colHasZero := make([]bool, n) // 列是否包含 0
	//[false,true,false]
	/**
	false
	true
	false
	*/
	//x=1,y=1
	for i, row := range matrix {
		for j, x := range row {
			if x == 0 {
				rowHasZero[i] = true
				colHasZero[j] = true
			}
		}
	}

	//[false,true,false]
	/**
	false
	true
	false
	*/
	//x=1,y=1
	for i, row0 := range rowHasZero {
		for j, col0 := range colHasZero {
			if row0 || col0 { // i 行或 j 列有 0
				matrix[i][j] = 0 // 题目要求原地修改，无返回值
			}
		}
	}
}
