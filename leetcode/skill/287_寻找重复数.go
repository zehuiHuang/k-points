package skill

/*
*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：

输入：nums = [1,3,4,2,2]
输出：2
示例 2：

输入：nums = [3,1,3,4,2]
输出：3
示例 3 :

输入：nums = [3,3,3,3,3]
输出：3
*/

//思路:做映射
/**
3,4,2,2
1->3
2->4
3->2
4->2
*/
func findDuplicate(nums []int) int {
	//快慢指针
	slow, fast := 0, 0
	//先模拟环形节点,找到相遇的点
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	//将慢指针移动到头部,快慢指针同时移动一步,相遇的点即为公共节点
	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	return slow
}
