package practice

/*
*
思路:通过队列结构进行辅助
1、将数据放入队列中
2、定位到k位置,然后将1到k-1的数据在重新塞入到队列头部
3、将找到的k从对立中移除,然后重复执行上一步,直到队列中数据只剩下一个
*/
func findTheWinner(n int, k int) int {
	queue := []int{}
	for i := 0; i < n; i++ {
		queue = append(queue, i)
	}
	for len(queue) > 1 {
		//因为索引是从0开始的所有是k-1
		v := (k - 1) % len(queue)
		q1 := queue[v+1:]
		q2 := queue[:v]
		if len(q1) == 0 {
			queue = q2
		} else if len(q2) == 0 {
			queue = q1
		} else {
			queue = append(q1, q2...)
		}
	}
	return queue[0] + 1
}
