package practice

/*
*
思路:回溯算法
1、定义left和right,left是已使用的左括号数量,right是已使用的右括号数量
2、并且要求left>right时才能选")" ,因为必须保证括号的顺序正确性,比如()) 就不对了
3、定义dfs函数:dfs(int,int,path)
4、终止条件:left==n && right==n
*/
func generateParenthesis(n int) []string {
	ans := make([]string, 0)

	var dfs func(int, int, string)
	dfs = func(left, right int, path string) {
		//终止条件,因为左括号和右括号数量一致,所以i==n时,说明已经生成了一个完整的括号组合
		if left == n && right == n {
			ans = append(ans, path)
			return
		}
		//选"(",
		if left < n {
			dfs(left+1, right, path+"(")
		}
		//选")",只有"("的数量大于"("的数量时,才能选
		if right < left {
			dfs(left, right+1, path+")")
		}
	}
	dfs(0, 0, "")
	return ans
}
