package practice

/*
*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
示例 2：

输入：n = 1, k = 1
输出：[[1]]
*/

/*
*组合数:关键index去掉重复
 */
func combine(n int, k int) [][]int {
	var dfs func(index int, path []int)
	ans := [][]int{}
	dfs = func(index int, path []int) {
		//条件判断
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		//重点理解:for循环是横向遍历,递归是纵向遍历
		//横向遍历是为了从1开始取值,比如取1,取2
		//纵向遍历是为了在取1的基础上,再取下一个,组合并判断是否符合条件

		//横向遍历,index就是为了重复,后续的递归只能选择该输的后面一个
		for i := index; i <= n; i++ {
			//纵向遍历,寻找符合条件的组合
			dfs(i+1, append(path, i))
		}
	}
	dfs(1, []int{})
	return ans
}
