package practice

import "sort"

/*
*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]
*/

/*
*
思路:回溯,和39题的区别有两点,一个是有重复数,一个不能重复选择
1、有重复数的思路是:先排序,对相邻的如果选过的就不能再选了
2、不能重复选择,那么就是下个迭代要startIndex+2
*/
func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	n := len(candidates)
	var dfs func(sum, index int, path []int)
	dfs = func(sum, index int, path []int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		//index为从哪个位置开始选择,因为选过的还可以再选(但组合不能重复),所有下一层迭代还是index,只有选择完不能再选了下层迭代才是index+1
		for i := index; i < n; i++ {
			if sum+candidates[i] > target {
				continue
			}
			//防止重复
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			//i+1是因为选过的数不能再选了
			dfs(sum+candidates[i], i+1, append(path, candidates[i]))
		}
	}
	dfs(0, 0, []int{})
	return ans
}
