package practice

/*
*
思路:递归算法
1、递归的思路,将数组分为左右两边,中间的数作为根节点,递归处理左右两边的数组
2、递归的边界条件,数组为空,则返回nil
3、递归的返回值,返回根节点
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := len(nums) >> 1
	return &TreeNode{
		Val:   nums[m],
		Left:  sortedArrayToBST(nums[:m]),
		Right: sortedArrayToBST(nums[m+1:]),
	}
}
