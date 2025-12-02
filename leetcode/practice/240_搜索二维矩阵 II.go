package practice

/*
*
思路:类似二叉搜索树
1、以左下角开始朝上遍历
2、如果当前位置>target,那么他的值一定在横坐标的上方,即i--,
3、如果当前值<target,那么target的值一定当前值的右边,即y++
4、如果前值==target,则直接返回
*/
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix) - 1
	n := len(matrix[0]) - 1
	i := m
	j := 0
	for i >= 0 && j <= n {
		if matrix[i][j] > target {
			//结果在它的左上方,x朝左上移动
			i--
		} else if matrix[i][j] < target {
			//结果在它的右上方,y右上方移动
			j++
		} else {
			return true
		}
	}
	return false
}
