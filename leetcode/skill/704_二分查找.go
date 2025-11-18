package skill

// 二分查找
func search(nums []int, target int) int {
	//排序-已经排序好了(升序),直接处理即可
	//0,1,2,3,4,5,6
	//0,1,2,3
	//0,1,2
	//5,6
	left, right := 0, len(nums)-1
	for left <= right {
		//m := (left + right) / 2
		/**
		left+(right-left)>>1==left + (right-left)/2==(2*left+right-left)/2==(left + right) / 2
		*/
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
