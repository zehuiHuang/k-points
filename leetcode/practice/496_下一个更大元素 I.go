package practice

/*
*
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

示例 1：

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：

输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
*/

/*
*思路:整体上是使用单调栈,将nums2用到单调栈
 */
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		ans[i] = -1
		//将nums1的下标和值做映射,方便从nums2的值对应到nums1上的下标
		mp[nums1[i]] = i
	}
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		//当 当前值大于栈顶元素,说明当前值在nums1中找到了第一个比它大的值,需放入ans
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			//栈顶元素
			index := stack[len(stack)-1]
			//栈顶元素对应的值
			value := nums2[index]
			//nmp[value]为栈顶元素对应值在nums1中的下标
			if _, ok := mp[value]; ok {
				ans[mp[value]] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
