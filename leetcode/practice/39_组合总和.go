package practice

import "sort"

/**
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。

示例 1：

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
示例 2：

输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
示例 3：

输入: candidates = [2], target = 1
输出: []

*/

/*
*
思路:回溯
1、组合问题都要有一个startIndex,是为了解决重复问题
2、如果是任意选择下一个迭代要startIndex+1,否则下一个迭代直接传旨startIndex即可
3、常规的终止条件判断和数据收集
4、最后别忘了要调用下dfs函数
*/
func combinationSum(candidates []int, target int) [][]int {
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
			//path = append(path, candidates[i])
			dfs(sum+candidates[i], i, append(path, candidates[i]))
			//因为都是复制,所有不需要回溯原来的值
			//path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []int{})
	return ans
}
