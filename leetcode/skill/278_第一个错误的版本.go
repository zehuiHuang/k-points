package skill

func firstBadVersion(n int) int {
	//二分查找
	//1,2,3,4,5,6,7,A,B,C mid:=5
	//6,7,A,B,C    mid:=A
	//6,7 mid:=6
	//7 mid:=
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		//错误的,则应该左移right
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	//left=A,是第一个错误版本
	return left
}

func isBadVersion(target int) bool {
	if target == 5 {
		return true
	}
	return false
}
