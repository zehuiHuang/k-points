package practice

import "sort"

/*
*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

/*
思路：
1. 先排序
2. 遍历所有区间，如果当前区间的起始位置（L）大于结果数组的最后一个区间的结束位置，则直接添加当前区间
3. 否则合并区间，取较大的结束位置
*/
func merge(intervals [][]int) [][]int {
	// 如果输入为空，直接返回空数组
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组
	merged := make([][]int, 0)

	// 遍历所有区间
	for i := 0; i < len(intervals); i++ {
		L, R := intervals[i][0], intervals[i][1]

		// 如果结果数组为空，或者当前区间与最后一个区间不重叠
		if len(merged) == 0 || merged[len(merged)-1][1] < L {
			// 直接添加当前区间
			merged = append(merged, []int{L, R})
		} else {
			// 否则合并区间，取较大的结束位置
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], R)
		}
	}

	return merged
}

// 思路:
// 先按照左区间进行生序排列,在判断第一个的右闭区间和第二个的左右区间做对比,
// 如果第一个右闭区间小于第二个左闭区间,则不能合并,如果大于等于,则还需要和第二个右闭区间对比,并取最大值进行合并
func merge2(intervals [][]int) [][]int {
	//生序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//合并
	ans := [][]int{}
	first := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if first[1] < intervals[i][0] {
			ans = append(ans, first)
			first = intervals[i]
		} else {
			if first[1] < intervals[i][1] {
				first[1] = intervals[i][1]
			}
		}
	}
	ans = append(ans, first)
	return ans
}
