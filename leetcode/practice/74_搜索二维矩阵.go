package practice

/*
*
思路:二分查找
1、其实可以将二维数组所有的多行数拼接成一维的有序数组
2、也可以在二位数组中采用二分查找
3、[i/n]便是所在行数:例如:m=3,n=4 ,i=9 ,那么在9/4=2 行
4、[i%n]便是所在列数:例如m=3,n=4 ,i=9 ,那么在9%4=1 列,可以简单理解为余数就是所在行的第几个位置
*/
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		x := matrix[mid/n][mid%n]
		if x == target {
			return true
		} else if target < x {
			right = mid + 1
		} else {
			left = mid + 1
		}
	}
	return false
}
