package practice

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
		//横向遍历
		for i := index; i <= n; i++ {
			//纵向遍历,寻找符合条件的组合
			dfs(i+1, append(path, i))
		}
	}
	dfs(1, []int{})
	return ans
}
