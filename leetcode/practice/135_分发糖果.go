package practice

/*
*
思路:贪心算法
1、先从左边开始,只判定右边i的是否比左边i-1的大,如果大那么i的小孩就会比i-1多个一个,依次遍历
2、再从右边开始,判定左边i小孩的糖果是否比右边i+1小孩糖果多,如果多那么 i的小孩就比i+1多一个,依次遍历
3、在第二步的遍历过程中,在同一个i位置的小孩,从分别从左边和右边计算出了可得到的糖果数,然后取两值的最大值
为什么取两值的最大值呢,因为针对i小孩的视角,第一次是i和i-1的比较得出应该要的结果,第二次是i和i+1比较的结果,取最大值才能都满足
*/
func candy(ratings []int) int {

	lArr := make([]int, len(ratings))
	for i, _ := range lArr {
		lArr[i] = 1
	}
	rArr := make([]int, len(ratings))
	for i, _ := range rArr {
		rArr[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			lArr[i] = lArr[i-1] + 1
		}
	}
	//因为在第二个遍历中还没有加上第一次遍历后的结果
	ans := lArr[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			rArr[i] = rArr[i+1] + 1
		}
		ans += max(lArr[i], rArr[i])
	}
	return ans
}
