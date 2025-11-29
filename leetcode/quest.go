package leetcode

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// 100, 4, 200, 1, 3, 2
//数组最长连续数

func LongestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	//连续长度
	longestStreak := 0
	for num := range numSet {
		//当前数字的前一个数字不存在，说明当前数字是一个连续序列的起点
		if !numSet[num-1] {
			//当前数字作为起点，开始查找连续序列
			currentNum := num
			//当前连续序列的长度
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			//如果当前连续序列的长度大于之前的最大长度，更新最大长度
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}

// 跳跃游戏leetcode:45
// 2, 3, 1, 1, 4  可跳步数
// 0  1  2  3  4  地址索引
/**
i+nums[i] 表示为下次最远跳跃 “下标位置”
原理：找到当前节点可跳的最远距离，然后计算最远距离到当前节点之间所有节点的可跳最远距离，找到最大值，即为要跳的位置，然后统计跳跃次数
*/
func jump(nums []int) int {
	//数组长度
	length := len(nums)
	//下次的最右起跳点
	end := 0
	//目前能跳到的最远位置
	maxPosition := 0
	//跳跃次数
	steps := 0
	for i := 0; i < length-1; i++ {
		//找到下一步最远可跳跃的距离
		maxPosition = max(maxPosition, i+nums[i])
		// 到达上次跳跃能到达的右边界了，就需要在右边界内找下次起跳的最远位置，同时步数加1
		if i == end {
			end = maxPosition // 目前能跳到的最远位置变成了下次起跳位置的右边界
			steps++           // 进入下一次跳跃
		}
	}
	return steps
}

// 获取两数最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 跳跃游戏leetcode:55
// 2, 3, 1, 1, 4  可跳步数案例1
// 0  1  2  3  4  地址索引
// 3, 2, 1, 0, 4  可跳步数案例2
// 0  1  2  3  4  地址索引

// 2 3 1 4 0 5
// 0 1 2 3 4 5
/**
原理：首先找到当前节点可跳的最远位置，那么最远位置的前面几个位置都是可以跳的，
那么遍历这几个可跳的点的下一个跳转位置A，判断这个A是否大于等于数组最大下标（遍历时，顺便把最大可跳的下标更新）
若不满足条件，则继续遍历（即找下一步可跳的最大位置）
*/
func canJump(nums []int) bool {
	n := len(nums) //长度
	rightmost := 0 //当前可跳的最远的下标
	for i := 0; i < n; i++ {
		if i <= rightmost { //当前位置小于等于最远位置，则可以跳到，否则无法跳了
			rightmost = max(rightmost, i+nums[i])
			if rightmost >= n-1 {
				return true
			}
		}
	}
	return false
}

// 加油站leetcode:134
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
/**
思路：首先遍历环型数组，如果一路总的剩余油量小于0，那么一定不能行驶一周，另外每循环一层时，需要计算从起点到当前节点的总耗油量，
如果小于零那么从i开始一定不能行驶一周
*/
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	//起点
	start := 0
	//一路总的剩余油量
	totalCount := 0
	//从起点开始的当前剩余油量
	currentCount := 0
	for i := 0; i < n; i++ {
		totalCount += gas[i] - cost[i]
		currentCount += gas[i] - cost[i]
		//如果当前剩余油量小于0，说明从start到i这段路无法走完，所以从i+1开始重新计算
		if currentCount < 0 {
			start = i + 1
			currentCount = 0
		}
	}
	//如果总的剩余油量小于0，说明无法走完一圈
	if totalCount < 0 {
		return -1
	}
	return start
}

/*
*
数组中的第K个最大元素；leetcode：215
*/
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

/*
*
l:左边界下标;r:右边界下标;k:第k个最大元素
输入: [3,2,1,5,6,4], k = 2
输出: 5
*/
func quickselect(nums []int, l, r, k int) int {
	//如果只有一个元素，则直接返回
	if l == r {
		return nums[k]
	}
	//选择左边界的元素作为分区点
	partition := nums[l]
	//左边界朝左移动一位
	i := l - 1
	//右边界朝右移动一位
	j := r + 1
	//循环，如果左边界朝左移动一位 小于 右边界朝右移动一位，则循环

	for i < j {
		//从最左边找，找到不小于分区点的值下标
		for i++; nums[i] < partition; i++ {
			fmt.Println(i)
		}
		//从最右边找，找到不大于分区点的值下标
		for j--; nums[j] > partition; j-- {
			fmt.Println(j)
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

/*
*
数组中的第K个最大元素；leetcode：215
一句话解释：首先通过快排，找到基准值所在的位置，保证基准值左边的都比它大，右边的都比它小，那么如果基准值的下标值+1等于k，那么就说明基准值就是第k大的元素
如果不是，那么继续在左边或者右边通过快排进行查找
*/
func findKthLargest2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	//pivotIndex的左边都比pivot大，右边都比pivot小
	pivotIndex := partition(nums, left, right)
	currentK := pivotIndex - left + 1 // 当前子数组中，基准值是第 currentK 大的元素
	// 如果当前基准值就是第 k 大的元素，直接返回
	if currentK == k {
		return nums[pivotIndex]
	} else if currentK > k {
		// 在左半部分寻找第 k 大的元素
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		// 在右半部分寻找第 (k - currentK) 大的元素
		return quickSelect(nums, pivotIndex+1, right, k-currentK)
	}
}

func partition(nums []int, left, right int) int {
	// 随机选择基准值，避免最坏情况
	randomIndex := rand.Intn(right-left+1) + left
	//将基准值放到最右边
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]
	pivot := nums[right]

	i := left - 1 // 记录大于基准的元素的边界
	for j := left; j < right; j++ {
		if nums[j] > pivot { // 将大于基准的元素移到左侧
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 将基准值放到正确位置（i+1）
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

/*
*
有序数组交集实现方法
方法：
*/
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		n1, n2 := nums1[i], nums2[j]
		if n1 == n2 {
			result = append(result, n1)
			i++
			j++
		} else if n1 < n2 {
			i++
		} else {
			j++
		}
	}
	return result
}

/*
*
leetcode:88：合并两个有序数组
双指针
*/
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	sorted := make([]int, 0, m+n)
	//p1执行nums1的指针，p2执行nums2的指针
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	//将排序后的数组复制到nums1，过程当中是覆盖
	return sorted
}

/*
*
双指针 数组逆序
*/
func reverseInPlace(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

/*
*
leetcode 49:字母异位词分组
*/
func groupAnagrams(strs []string) [][]string {
	//26个字母,设置map，key为26个字母的数组，value为字符串数组
	mp := map[[26]int][]string{}
	//循环遍历数组中的字符串
	for _, str := range strs {
		cnt := [26]int{}
		//循环遍历字符串中的字符
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

/*
*
560. 和为 K 的子数组
描述：
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。
*/
func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

/*
leetcode:560 和为 K 的子数组
描述：
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

输入：nums = [1,2,3], k = 3
输出：2
思路：
1. 利用map来存储前缀和的出现次数
2、如果前缀和减去k的值在map中存在，则说明存在一个子数组，其和为k
*/
func subarraySum2(nums []int, k int) int {
	count, pre := 0, 0
	//key为前缀和，value为前缀和出现的次数
	m := map[int]int{}
	//前缀和为0的初始化数据为1次
	m[0] = 1
	//循环数组
	for i := 0; i < len(nums); i++ {
		//计算前缀和
		pre += nums[i]
		//如果前缀和减去k的值在map中存在，则说明存在一个子数组，其和为k
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		//防止值为0的出现导致少算，所以是+=1
		m[pre] += 1
	}
	return count
}

// leetcode：56 合并区间
// merge2 合并重叠的区间
// 参数：intervals - 区间数组，每个区间是一个长度为2的数组，表示区间的起始和结束
// 返回：合并后的区间数组
/**
思路：
1. 先排序
2. 遍历所有区间，如果当前区间的起始位置（L）大于结果数组的最后一个区间的结束位置，则直接添加当前区间
3. 否则合并区间，取较大的结束位置
*/
func merge2(intervals [][]int) [][]int {
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

/*
*

	leetcode:189 旋转数组

描述：
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
*/
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func abc(nums []int) int {
	ans := math.MinInt64
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			ans = max(ans, count)
		}
	}
	return ans
}

// 连续子数组的最大和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentMax := nums[0]
	globalMax := nums[0]
	//-2, 11, -4, 15, -5, -2
	for i := 1; i < len(nums); i++ {
		// 决定是延续当前子数组还是从当前元素重新开始
		currentMax = max(nums[i], currentMax+nums[i])
		// 更新全局最大值
		globalMax = max(globalMax, currentMax)
	}
	return globalMax
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // 交换字符位置
	}
	return string(runes)
}

func largestNumber(nums []int) string {
	tmp := make([]string, len(nums))
	for i := range nums {
		tmp = append(tmp, strconv.Itoa(nums[i]))
	}
	sort.Slice(tmp, func(i, j int) bool {
		v1 := tmp[i] + tmp[j]
		v2 := tmp[j] + tmp[i]
		return v1 > v2
	})
	ans := strings.Builder{}
	for i := 0; i < len(tmp); i++ {
		ans.WriteString(tmp[i])
	}
	return ans.String()
}
