package skill

// 方法1
func rotate(matrix [][]int) {
	//先水平翻转,在对角线翻转
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		//水平翻转的规律:x坐标:i->n-1-i。y坐标:n-1-i->i
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	//对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// 方法2
func rotate2(matrix [][]int) {
	/**
	[5  1  9   11]
	[2  4  8   10]
	[13 3  6   7]
	[15 14 12  16]
	*/
	//旋转90度,找规律规律:对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。

}
