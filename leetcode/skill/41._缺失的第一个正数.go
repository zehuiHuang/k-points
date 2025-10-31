package skill

// 核心思想是: 对数组为nums(长度为N)中未出现的最小正整数,只能在1~N的范围内,如果都不在,那么最小的一定是N+1
// 第一步:将所有负数都设置成比N大的即可(N+1)
// 第二步:打标记,将每个小于N的数x, 将数组的x-1位置的数都标记为负数(该处是最不好理解的,这里最终是为了统计数组的位置,最终求也是1~N的位置)
// 第三步:第一个不是负数的位置,即为答案
func firstMissingPositive(nums []int) int {
	N := len(nums)
	//1设置成正整数
	for i := 0; i < N; i++ {
		if nums[i] <= 0 {
			nums[i] = N + 1
		}
	}
	//[3,4,-1,1] -》[3,4,5,1]

	//2打标记
	for i := 0; i < N; i++ {
		//在这个范围内的
		num := abs(nums[i])
		if num <= N {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	//[3,4,4,1] -》
	//[3,4,4,1]

	res := N + 1
	for i := 0; i < N; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
